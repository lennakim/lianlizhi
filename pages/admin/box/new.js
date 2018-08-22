const qiniuUploader = require("../../../utils/qiniuUploader");

function initQiniu(uptoken) {

  var options = {
    region: 'SCN',
    uptoken: uptoken,
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

    var uptoken = null;

    wx.request({
      method: "POST",
      url: "http://47.98.157.166:2930/qiniu/uptoken/",
      success: function(res) {
        uptoken = res.data.data
        initQiniu(uptoken) //
      }
    })


    var self = this;
    wx.chooseImage({
      count: 1, // 默认9

      sourceType: ['album', 'camera'], // 可以指定来源是相册还是相机，默认二者都有
      success: function(res) {
        var filePath = res.tempFilePaths[0];

        qiniuUploader.upload(filePath, function(res) {
          self.setData({
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