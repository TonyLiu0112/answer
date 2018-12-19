package handlers

import (
	"com.tony666.answer/activity"
	"net/http"
)

var c = []activity.Card{
	{
		activity.Question{Id: 1, Title: "第一题", Content: "中国最大的淡水湖是?", Options: []activity.Option{{1, 1, "A: 鄱阳湖"}, {1, 2, "B: 长江"}}},
		activity.Explanation{Content: "鄱阳湖的古称很多：彭蠡泽、彭泽、官亭湖、扬澜、担石湖等等，不下10个。这倒不是因为它有许多渊源，而是由于它兼并了许多小湖，逐渐扩大，同时也并蓄了那些小湖的名字。它本初的乳名源自大禹治水时期，这片地区因地势低洼，形成了数条分汊状水系，所以取古汉语中表数量多的虚词“九”，称其为九江。《禹贡》中记载“九江孔殷，东为彭蠡。”彭者，大也；蠡者，瓠瓢也，即葫芦。也就是说，这片洼地湖泊，自古就形似葫芦瓢。"},
		activity.Solution{Key: 1},
	},
	{
		activity.Question{Id: 2, Title: "第二题", Content: "中国最高的山是?", Options: []activity.Option{{2, 1, "A: 泰山"}, {2, 2, "B: 珠穆朗玛峰"}}},
		activity.Explanation{Content: "珠穆朗玛峰（珠峰）是喜马拉雅山脉的主峰，是世界海拔最高的山峰，位于中国与尼泊尔边境线上，它的北部在中国西藏定日县境内（西坡在定日县扎西宗乡，东坡在定日县曲当乡，有珠峰大本营），南部在尼泊尔境内，而顶峰位于中国境内，是世界最高峰。是中国跨越四个县的珠穆朗玛峰自然保护区和尼泊尔国家公园的中心所在。"},
		activity.Solution{Key: 2},
	},
	{
		activity.Question{Id: 3, Title: "第三题", Content: "16-17赛季欧冠联赛冠军是?", Options: []activity.Option{{3, 1, "A: 马德里竞技"}, {3, 2, "B: 皇家马德里"}}},
		activity.Explanation{Content: "2016年5月29日2时45分，第61届冠军杯也是第24届冠军联赛决赛在米兰圣西罗球场打响，皇家马德里足球俱乐部在120分钟内与马竞战为1比1平，点球战以5比3胜出，创纪录第11次夺冠，马竞成为首支前3届冠军杯决赛全部告负的球队。"},
		activity.Solution{Key: 2},
	},
	{
		activity.Question{Id: 4, Title: "第四题", Content: "三国中的庞统字什么?", Options: []activity.Option{{4, 1, "A: 卧龙"}, {4, 2, "B: 凤雏"}}},
		activity.Explanation{Content: "庞统（179年－214年），字士元，号凤雏，汉时荆州襄阳（治今湖北襄阳）人。东汉末年刘备帐下重要谋士，与诸葛亮同拜为军师中郎将。与刘备一同入川，于刘备与刘璋决裂之际，献上上中下三条计策，刘备用其中计。进围雒县时，庞统率众攻城，不幸中流矢而亡，年仅三十六岁，追赐统为关内侯，谥曰靖侯。后来庞统所葬之处遂名为落凤坡。"},
		activity.Solution{Key: 2},
	},
	{
		activity.Question{Id: 5, Title: "第五题", Content: "Switch是以下哪个公司的产品?", Options: []activity.Option{{5, 1, "A: 任天堂"}, {5, 2, "B: 索尼"}}},
		activity.Explanation{Content: "NS，全名NINTENDO SWITCH，是任天堂游戏公司于2017年3月首发的旗舰产品，主机采用家用机掌机一体化设计。新机不锁区，支持1920*1080电视输出和1280*720掌上输出。港版NS于2017年3月3日发售，台版NS于12月1日发售。 [1]  这是前社长岩田聪最后一部参与开发的硬件产品，该产品将成为未来任天堂娱乐事业蓝图的中心。NS首秀获得强烈反响，预告片YouTube首日播放量超一千万回，一度登顶YouTube播放榜首，风头压过美国大选。"},
		activity.Solution{Key: 1},
	},
}

func BeginActivity(w http.ResponseWriter, r *http.Request) {
	activity.CrtActivity.Build(c)
	go activity.CrtActivity.Begin()
	_, _ = w.Write([]byte("问答将在5秒后开始....."))
}
