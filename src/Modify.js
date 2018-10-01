import React, { PureComponent } from 'react';
import {
  NavBar,
  Toast,
  Icon,
  List,
  WingBlank,
  WhiteSpace,
  Button
} from 'antd-mobile';
import {
  employeeDetail,
  employeeRemove,
  employeeModify,
  employeeSignIn
} from './service/employee';
import { timestampToTime } from './utils';

const Item = List.Item;

class Modify extends PureComponent {
  state = {
    employee: {},
    loading: false
  };

  componentDidMount() {
    const {
      match: {
        params: { id }
      }
    } = this.props;
    this.loadDetail(id);
  }

  loadDetail = id => {
    Toast.loading('数据加载中...', 0);
    employeeDetail(id).then(res => {
      const { data } = res;
      if (data.code !== '0') {
        Toast.fail(data.message);
      } else {
        Toast.hide();
        const { employee } = data;
        this.setState({
          employee
        });
      }
    });
  };

  onRefreshToken = evt => {
    const {
      history: { push },
      match: {
        params: { id }
      }
    } = this.props;
    const {
      employee: { Mobile, Password }
    } = this.state;
    if (evt !== undefined && evt.preventDefault) evt.preventDefault();
    Toast.loading('数据更新中...', 0);
    employeeModify(id, { mobile: Mobile, password: Password }).then(res => {
      const { data } = res;
      if (data.code !== '0') {
        Toast.fail(data.message, 1, () => {
          push('/');
        });
      } else {
        Toast.success('token更新成功', 1, () => this.loadDetail(id));
      }
    });
  };

  onSignClick = evt => {
    const {
      history: { push },
      match: {
        params: { id }
      }
    } = this.props;
    if (evt !== undefined && evt.preventDefault) evt.preventDefault();
    Toast.loading('签到中...', 0);
    employeeSignIn(id).then(res => {
      const { data } = res;
      if (data.code !== '0') {
        Toast.fail(data.message, 1, () => {
          push('/');
        });
      } else {
        Toast.success('签到成功', 1, () => this.loadDetail(id));
      }
    });
  };

  onRemoveClick = evt => {
    const {
      history: { push },
      match: {
        params: { id }
      }
    } = this.props;
    if (evt !== undefined && evt.preventDefault) evt.preventDefault();
    Toast.loading('数据删除中...', 0);
    employeeRemove(id).then(res => {
      const { data } = res;
      if (data.code !== '0') {
        Toast.fail(data.message, 1, () => push('/'));
      } else {
        Toast.success(data.message, 1, () => push('/'));
      }
    });
  };

  render() {
    const {
      history: { goBack }
    } = this.props;
    const { employee, loading } = this.state;
    return (
      <div>
        <NavBar
          mode="dark"
          leftContent={<Icon type="left" />}
          onLeftClick={goBack}
        >
          {employee.Name}
        </NavBar>
        <List renderHeader={() => '姓名'}>
          <Item>{employee.Name}</Item>
        </List>
        <List renderHeader={() => '员工号'}>
          <Item>{employee.EmployeeID}</Item>
        </List>
        <List renderHeader={() => '手机号'}>
          <Item>{employee.Mobile}</Item>
        </List>
        <List renderHeader={() => 'Token（点击更新）'}>
          <Item
            arrow="horizontal"
            multipleLine
            onClick={evt => this.onRefreshToken(evt)}
          >
            {employee.UserToken}
          </Item>
        </List>
        <List renderHeader={() => '最近签到'}>
          <Item>
            {employee.LastSignInTime === 0
              ? '无签到记录'
              : timestampToTime(employee.LastSignInTime)}
          </Item>
        </List>
        <WingBlank>
          <WhiteSpace />
          <Button
            type="primary"
            loading={loading}
            onClick={evt => this.onSignClick(evt)}
          >
            签到
          </Button>
          <WhiteSpace />
          <Button
            type="warning"
            loading={loading}
            onClick={evt => this.onRemoveClick(evt)}
          >
            删除
          </Button>
        </WingBlank>
      </div>
    );
  }
}

export default Modify;
