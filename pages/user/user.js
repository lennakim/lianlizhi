Page({

  data: {
    checkout: ['微信', '支付宝', '银行转账']
  },

  bindPickerChange: function (e) {
    console.log('picker 数值是', e.detail.value)

    this.setData({
      inx: e.detail.value
    })
  },

  enterAdmin: function () {
    wx.navigateTo({
      url: '/pages/admin/login'

    })

    wx.setNavigationBarTitle({
      title: '后台'
    })
    
  }
})