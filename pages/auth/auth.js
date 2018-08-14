// pages/auth/auth.js
Page({

  /**
   * 页面的初始数据
   */
  data: {

  },

  /**
   * 生命周期函数--监听页面加载
   */
  onLoad: function(options) {

  },

  /**
   * 生命周期函数--监听页面初次渲染完成
   */
  onReady: function() {

  },

  /**
   * 生命周期函数--监听页面显示
   */
  onShow: function() {

  },

  /**
   * 生命周期函数--监听页面隐藏
   */
  onHide: function() {

  },

  /**
   * 生命周期函数--监听页面卸载
   */
  onUnload: function() {

  },


  onGotUserInfo: function(e) {
    if (undefined != e.detail['userInfo']) { // 允许授权
      wx.setStorageSync('userInfo', e.detail.userInfo)

      wx.navigateTo({
        url: "/pages/user/user"
      })

      // console.log(e.detail.rawData)
    } else { // 拒绝授权
      wx.showModal({
        title: '提示',
        content: '小程序需要您的授权才能提供更好的服务哦',
        showCancel: false,
        confirmText: "好的"
      })

    }


  },
})