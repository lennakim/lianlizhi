const qiniuUploader = require("../../../utils/qiniuUploader");

function initQiniu() {
  var options = {
    region: 'SCN',
    uptokenFunc: function() {
      return ""
    },
    domain: 'http://pdum902ru.bkt.clouddn.com/',
    shouldUseQiniuFileName: false
  };

  qiniuUploader.init(options);
}

Page({

  data: {
    imgObj: {}
  },
  onLoad: function(options) {

  },
  onReady: function() {

  },
  /**
   * 用户点击右上角分享
   */
  onShareAppMessage: function() {

  },
  chooseImage: function(e) {
    initQiniu();

    var that = this;
    wx.chooseImage({
      count: 1, // 默认9

      sourceType: ['album', 'camera'], // 可以指定来源是相册还是相机，默认二者都有
      success: function(res) {
        var filePath = res.tempFilePaths[0];

        qiniuUploader.upload(filePath, function(res) {
          that.setData({
            'imgObj': res // res.imageURL
          });
        })
      },
      error: function(err) {
        console.log('error: ' + JSON.stringify(err));
      }
    })
  },
  previewImage: function(e) {
    wx.previewImage({
      current: e.currentTarget.id, // 当前显示图片的http链接
      urls: this.imgObj.imageURL // 需要预览的图片http链接列表
    })
  }
})