import axios from 'axios';
import {message} from "ant-design-vue"
//设置请求超时
axios.defaults.timeout =  600000;
axios.defaults.baseURL = ''


/*
  封装get 方法
  @param url
  @param data
  @retrns {Promise}
*/

export function fetch(url,params ={}){
    let hide = this.$message.loading('加载中..', 0);
    return new Promise((resolve,reject) =>{
        axios.get(url,{
            params:params
        }).then(response =>{

            setTimeout(hide, 0);
            if (response.data.Code!=200){
                this.$message.error(response.data.Msg);
            }else{
                resolve(response.data)
            }
        })
            .catch(err =>{
                setTimeout(hide, 0);
                this.$message.error('加载失败,请重试');
                reject(err)
            })
    })
}


/*
   封装post 请求
   @param url
   @param data
   @returns {Promise}
*/

export function post(url,data ={}){
    let hide = this.$message.loading('加载中..', 0);
    return new Promise((resolve,reject) =>{
        axios.post(url,data).then(response =>{
            setTimeout(hide, 0);
            resolve(response.data)
        },err =>{
            setTimeout(hide, 0);
            reject(err)
        })
    })
}
