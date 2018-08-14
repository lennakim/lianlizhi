Page({

  data: {
    checkout: ['微信', '支付宝', '银行转账'],
    userInfo: wx.getStorageSync('userInfo'),
  },
  onLoad: function(options) {
    if (!wx.getStorageSync('userInfo')) {
      wx.navigateTo({
        url: "/pages/auth/auth"
      })
    }

  },
  bindPickerChange: function(e) {
    this.setData({
      inx: e.detail.value
    })
  },

  enterAdmin: function() {
    wx.navigateTo({
      url: '/pages/admin/login'

    })

    wx.setNavigationBarTitle({
      title: '后台'
    })

  }
})