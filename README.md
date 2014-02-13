# Credit Risk Modeling using Internet User Behavior Data

<table>
<tr>
<td width="50%">
In recent months, the idea of Internet Finance becomes so hot in China.  During the few days of 2014 Chinese New Year vacation, over 50 million deposit accounts were associated with [Weixin (a.k.a WeChat)](http://en.wikipedia.org/wiki/WeChat), a Chinese mobile social network app with over 300 million active users.  As Weixin supports payment in using the associated deposit accounts for taxi, restaurants and shopping, people can now throw away their wallets, deposit cards and credit card.  The next step is, very likely, finance management. For example, peer-to-peer lending, a high-returns Internet finance business.
</td>
<td>
最近互联网金融这个概念那是相当的火。刚刚过完的春节里，微信通过``红包''这个产品特性，吸引了很多人把银行卡关联到微信上。春节一过，微信就在与嘀嘀打车的合作之外，开辟了新的空间，让更多商家接收微信支付。
<br><br>
和支付相关的就是理财。而目前阿里理财宝、百度白发、微信理财都是用用户的钱去买货币基金。换句话说，只是传统金融产品的新的销售渠道。真正的互联网金融的价值应该是让钱作为资本流转得更高效。小额贷款是一个重要的路子。
</td>
<tr>
</table>

<table>
<tr>
<td width="50%">
At the heart of lending, it is the credit risk modeling. When we think about porting risk modeling to the Internet economy, it is nature to think about extracting more user information from Internet data and making the risk prediction more accurate.  
</td>
<td>
贷款的核心技术是风险预估模型。传统金融行业也有风险预估模型，但是因为用户信息有限，这些模型相对互联网行业（比如搜索和广告）里的模型就显得单薄了。在互联网行业，我们应该收集用户的互联网行为数据，获取更多用户信息，来更加精准的做风险预估。
</td>
</table>

<table>
<tr>
<td width="50%">
However, there is a difficulty here. Though we can collect data of Internet users and we do have credit records of borrowers, we do not know the mapping between Internet users and borrowers.
</td>
<td>
但是这里有一个问题。虽然互联网金融公司拥有贷款者的贷款和还款记录，也能收集很多互联网用户的信息，但是并不知道那些贷款者对应于哪些互联网用户。
</td>
</tr>
</table>

<table>
<tr>
<td width="50%">
In this following tutorial on Internet credit risk modeling, we introduce how to design a probabilistic model that learns not only risk prediction but also matching borrowers with Internet users whose behavior data had been collected.  
</td>
<td>
这篇tutorial就针对这个问题展开。我们考虑设计一个不同于传统金融行业里地风险预估模型，它除了能预估风险，还能把贷款者和互联网用户对应起来。
</td>
</tr>
</table>

<table>
<tr>
<td width="50%">
As a tutorial, the introduced model is not a sophisticated one; instead, it is something like a mixture logistic regression model.  It is not a industrially-deployable one either.  An deployable model requires considering more difficulties.  For example, Internet behavior data were collected from various products (search engine, e-shopping, and etc), and we have various kinds of Internet users -- search engine users, e-shopping users and etc.  Therefore, the matching would be between borrower and these kinds of Internet users.  Also, it is very important to incorporate prior knowledge about the matching into the model.  Maybe we will address these issue in further tutorials.
</td>
<td>
因为这只是一个tutorial，只是为了吸引大家对互联网金融中的技术问题的关注，所以我们介绍的模型不是一个超牛逼模型——实际上看起来只是一个mixture of logistic regression模型。当然，沿着这个tutorial的思路，你可以很容易的把它扩展成更牛逼的模型。另外，这个模型也不足以直接用在业务里，因为还有更多的问题需要考虑。比如互联网用户是包括很多种类的，比如百度用户、微博用户等等。我们需要知道每个贷款者对应哪个百度用户、哪个微博用户等等。而且，一个可以工业应用的模型还应该考虑如何引入先验知识，描述大致已知的贷款者和互联网用户的对应关系。如果将来还有时间，可以考虑把这篇tutorial补充补充，加上对这些问题的考虑。
</td>
</tr>
</table>

<table>
<tr>
<td width="50%">
Currently, the tutorial is in Chinese: <a href=""https://github.com/wangkuiyi/risk_model/blob/master/tutorial/tutorial-cn.pdf">download here</a>
</td>
<td>
目前这个tutorial只有中文版。可以点<a href=""https://github.com/wangkuiyi/risk_model/blob/master/tutorial/tutorial-cn.pdf">这里</a>下载。
</td>
</tr>
</table>

<table>
<tr>
<td width="50%">
I also implement the risk model introduced in this tutorial using the Go language.  You can check out and run the demo program using the following commands:
</td>
<td>
我用Go语言写了这个模型的实现。可以执行下面的命令来下载源代码并且编译执行之。
</td>
</tr>
</table>

    export GOPATH=`pwd`
    go get github.com/wangkuiyi/risk_model
    ./bin/risk_model

