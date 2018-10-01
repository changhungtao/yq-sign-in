import request from './request';

// 创建员工
export const employeeCreate = params => request('/employee', params, 'post');

// 删除员工
export const employeeRemove = id => request(`/employee/${id}`, {}, 'delete');

// 编辑员工
export const employeeModify = (id, params) =>
  request(`/employee/${id}`, params, 'put');

// 员工列表
export const employeeList = params => request('/employee', params);

// 员工详情
export const employeeDetail = id => request(`/employee/${id}`);

// 员工签到
export const employeeSignIn = id => request(`/employee/${id}/sign`, {}, 'post');
