var boxData = [
  {
    "title": "克洛伊",
    "price": "2300",
    "img": "/images/chloe.jpg",
    "desc": "米白色女士手提双肩包",
    "state": "上架"
  },
  {
    "title": "香奈儿",
    "price": "6000",
    "img": "/images/lv.jpg",
    "desc": "香奈儿焦糖色牛皮女士单肩斜挎包",
    "state": "上架"
  },
  {
    "title": "迪奥",
    "price": "6832",
    "img": "/images/dior.jpg",
    "desc": "迷你牛皮斜挎包",
    "state": "上架"
  }
]


var checkoutBox = [
  {
    "title": "香奈儿",
    "price": "6000",
    "img": "/images/lv.jpg",
    "desc": "香奈儿焦糖色牛皮女士单肩斜挎包",
    "state": "已出售"
  },
  {
    "title": "克洛伊",
    "price": "2300",
    "img": "/images/chloe.jpg",
    "desc": "米白色女士手提双肩包",
    "state": "已出售"
  }
]


var expireBox = [
  {
    "title": "克洛伊",
    "price": "2300",
    "img": "/images/chloe.jpg",
    "desc": "米白色女士手提双肩包",
    "state": "过期"
  }
]

const sliderWidth = 96; // 需要设置slider的宽度，用于计算中间位置

Page({
  data: {
    tabs: ["寄卖列表", "结账列表", "到期列表"],
    activeIndex: 1, // 显示哪个 navbar
    sliderOffset: 0,
    sliderLeft: 0,
    boxs: boxData,
    checkouts: checkoutBox,
    expires: expireBox,
  },
  onLoad: function () {
    var self = this;

    wx.getSystemInfo({
      success: function (res) {
        self.setData({
          sliderLeft: (res.windowWidth / self.data.tabs.length - sliderWidth) / 2,
          sliderOffset: res.windowWidth / self.data.tabs.length * self.data.activeIndex
        });
      }
    });
  },
  tabClick: function (e) {
    this.setData({
      sliderOffset: e.currentTarget.offsetLeft,
      activeIndex: e.currentTarget.id
    });
  }
});