import axios from 'axios';

const baseURL = '';
const port = '';
const URL = port ? `${baseURL}:${port}/api` : `${baseURL}/api`;

export default (url, options = {}, method = 'get') => {
  const key = ~['get', 'head'].indexOf(method) ? 'params' : 'data';
  return axios({
    url: URL + url,
    method,
    validateStatus: false,
    ...{ [key]: options }
  }).then(res => res);
};
