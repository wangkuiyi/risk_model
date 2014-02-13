package main

import (
	"fmt"
	hector "github.com/xlvector/hector/core"
	hector "github.com/xlvector/hector/lr"
	"math"
	"math/rand"
)

type creditRecord struct {
	borrower string // The borrower ID, zero-based.
	borrowed int    // The number of times he borrowed.
	returned int    // The number of borrows he returned.
}

type internetUser []string // Every Internet user is represented by a set of binary features.

func main() {
	records := []creditRecord{
		{"Alice", 10, 9},
		{"Bob", 50, 1},
	}
	iusers := []internetUser{
		{"male", "30"},
		{"female", "20"},
	}
	match := [][]float64{
		{0.5, 0.5}, // Alice
		{0.5, 0.5}, // Bob
	}

	lr, id2f := train(records, iusers, match)
	printModel(lr, id2f)
	fmt.Printf("%v\n", match)
}

func train(records []creditRecord, iusers []internetUser, match [][]float64) (*lr.LogisticRegression, []string) {
	lr := new(lr.LogisticRegression)
	protos, id2f := constructFeatureVectors(iusers)
	borrower2iuser := make([]int, len(records))

	for iter := 0; iter < 100; iter++ {
		// M-step:
		sampleBorrower2IUser(match, borrower2iuser)
		dataset := constructTrainingData(records, borrower2iuser, protos)
		lr.Init(map[string]string{"learning-rate": "0.1", "regularization": "1.0", "steps": "20"})
		lr.Train(dataset)

		printModel(lr, id2f)	// debug

		// E-step:
		updateMatch(lr, records, protos, match)
	}

	return lr, id2f
}

func constructFeatureVectors(iusers []internetUser) ([]*core.Sample, []string) {
	protos := make([]*core.Sample, len(iusers))
	f2id := make(map[string]int)
	id2f := make([]string, 0)

	for i, u := range iusers {
		protos[i] = core.NewSample()

		for _, f := range u {
			id, exists := f2id[f]
			if !exists {
				id = len(id2f)
				id2f = append(id2f, f)
				f2id[f] = id
			}
			protos[i].AddFeature(core.Feature{int64(id), 1.0})
		}
	}
	return protos, id2f
}

func constructTrainingData(records []creditRecord, borrower2iuser []int, protos []*core.Sample) *core.DataSet {
	data := core.NewDataSet()
	for borrower, record := range records {
		for i := 0; i < record.borrowed; i++ {
			s := protos[borrower2iuser[borrower]].Clone()
			if i < record.returned {
				s.Label = 1
			} else {
				s.Label = 0
			}
			data.AddSample(s)
		}
	}
	return data
}

func sampleBorrower2IUser(match [][]float64, borrower2iuser []int) {
	for borrower, dist := range match {
		borrower2iuser[borrower] = cumulativeSample(dist)
	}
}

func cumulativeSample(dist []float64) int {
	choice := sum(dist) * rand.Float64()
	sum_so_far := 0.0
	for i, p := range dist {
		sum_so_far += p
		if sum_so_far >= choice {
			return i
		}
	}
	return -1
}

func sum(dist []float64) float64 {
	sum := 0.0
	for _, p := range dist {
		sum += p
	}
	return sum
}

func updateMatch(lr *lr.LogisticRegression, records []creditRecord, protos []*core.Sample, match [][]float64) {
	predictions := make([]float64, len(protos))
	for i, proto := range protos {
		predictions[i] = lr.Predict(proto)
	}

	for borrower, dist := range match {
		r := records[borrower].returned
		nr := records[borrower].borrowed - r

		for iuser, gamma := range dist {
			match[borrower][iuser] = gamma *
				math.Exp(float64(r)*math.Log(1-predictions[iuser]) + float64(nr)*math.Log(predictions[iuser]))
		}

		norm := sum(match[borrower])
		for iuser, prob := range match[borrower] {
			match[borrower][iuser] = prob / norm
		}
	}
}

func printModel(lr *lr.LogisticRegression, id2f []string) {
	for f, g := range lr.Model {
		fmt.Printf("%s : %f\n", id2f[f], g)
	}
}
