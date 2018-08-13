//index.js
//获取应用实例
const app = getApp()

var boxData = [
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

Page({
  data: {
    boxs: boxData,
    canIUse: wx.canIUse('button.open-type.getUserInfo')
  },
  //事件处理函数
  bindViewTap: function () {
    wx.navigateTo({
      url: '../logs/logs'
    })
  },

  onLoad: function () {

  },

  alertFun: function (e) {
    wx.showToast({
      title: '成功',
      icon: 'success'
    })

  }
})
