
function zeroPadding(num) {
  return ('00' + num).slice(-2)
}

let now = new Vue({
  el: '#now',
  data: {
    now: ''
  },
  methods: {
    updateNow: function (date) {
      this.now = date
    }
  },
  mounted: function () {
    let loc = window.location;
    let uri = 'ws:';

    if (loc.protocol === 'https:') {
      uri = 'wss:';
    }
    uri += '//' + loc.host;
    uri += loc.pathname + 'ws';

    let ws = new WebSocket(uri)

    ws.onopen = function () {
      console.log('Connected')
    }
    ws.onmessage = function (evt) {
      let date = new Date(evt.data * 1000);
      let year = date.getFullYear()
      let month = zeroPadding(date.getMonth() + 1)
      let day = zeroPadding(date.getDate())
      let hour = zeroPadding(date.getHours())
      let minutes = zeroPadding(date.getMinutes())
      let seconds = zeroPadding(date.getSeconds())
      let date_format = year + "/" + month + "/" + day + " " + hour + ":" + minutes + ":" + seconds
      now.updateNow(date_format)
    }
  }
})