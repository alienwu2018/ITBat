import axios from 'axios'



axios.defaults.withCredentials=true;


export default {
  post (option) {
    return axios({
      method: 'post',
      baseURL: '',
      url: option.url,
     // data: qs.stringify(option.data),
       data: option.data,
      timeout: 50000,
      headers: {
      //  'Content-Type': 'application/x-www-form-urlencoded;charset=UTF-8'
       'Content-Type': option.type
      }
    })
  },
  get (option) {
    return axios({
      method: 'get',
      baseURL: '',
      url: option.url,
      timeout: 50000,
      headers: {
        // 'Content-Type': type !== 'json' ? 'application/x-www-form-urlencoded' : 'application/json;charset=UTF-8'
        'Content-Type': 'application/json;charset=UTF-8'
      }
    })
  }
}