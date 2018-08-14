//app.js
App({
  onLaunch: function() {

    // 登录
    wx.login({
      success: function(res) {
        if (res.code) {
          // 发送 res.code 到后台换取 openId, sessionKey, unionId
          console.log('ok: ' + res.code)

        } else {
          console.log('fail: ' + res.errMsg)
        }
      }
    })

    // 获取用户信息
    wx.getSetting({
      success: res => {
        if (res.authSetting['scope.userInfo']) {
          wx.getUserInfo({
            success: res => {
              // 可以将 res 发送给后台解码出 unionId
              this.globalData.userInfo = res.userInfo
              wx.setStorageSync('userInfo', res.userInfo)

              // 由于 getUserInfo 是网络请求，可能会在 Page.onLoad 之后才返回
              // 所以此处加入 callback 以防止这种情况
              if (this.userInfoReadyCallback) {
                this.userInfoReadyCallback(res)
              }
            }
          })
        } else {
          console.log("用户信息没有授权")
        }
      }
    })
  },

  globalData: {
    userInfo: wx.getStorageSync('userInfo') || null,
    host: 'http://localhost:2930'
  }
})