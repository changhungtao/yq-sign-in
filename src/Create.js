import React, { PureComponent } from 'react';
import {
  NavBar,
  Toast,
  Icon,
  List,
  InputItem,
  WhiteSpace,
  WingBlank,
  Button
} from 'antd-mobile';
import MD5 from 'crypto-js/md5';

import { employeeCreate } from './service/employee';

class Create extends PureComponent {
  state = {
    hasMobileError: false,
    hasPwdError: false,
    mobile: '',
    pwd: '',
    loading: false
  };

  onMobileErrorClick = () => {
    if (this.state.hasMobileError) {
      Toast.info('请输入11位数字');
    }
  };

  onMobileChange = mobile => {
    if (mobile.replace(/\s/g, '').length < 11) {
      this.setState({
        hasMobileError: true
      });
    } else {
      this.setState({
        hasMobileError: false
      });
    }
    this.setState({
      mobile
    });
  };

  onPwdErrorClick = () => {
    if (this.state.hasPwdError) {
      Toast.info('密码不能为空');
    }
  };

  onPwdChange = pwd => {
    if (!pwd) {
      this.setState({
        hasPwdError: true
      });
    } else {
      this.setState({
        hasPwdError: false
      });
    }
    this.setState({
      pwd
    });
  };

  transformData = state => {
    const { mobile, pwd } = state;
    return {
      mobile: mobile.replace(/\s/g, ''),
      password: pwd ? MD5(pwd).toString() : ''
    };
  };

  onButtonClick = evt => {
    const {
      history: { push }
    } = this.props;
    const { mobile, pwd } = this.state;
    if (!mobile || !pwd) {
      Toast.fail('手机号或密码不能为空');
      return;
    }
    if (evt !== undefined && evt.preventDefault) evt.preventDefault();
    this.setState({
      loading: true
    });
    employeeCreate(this.transformData(this.state)).then(res => {
      this.setState({
        loading: false
      });
      const { data } = res;
      if (data.code !== "0") {
        Toast.fail(data.message);
      } else {
        Toast.success(data.message, 1, () => {
          push('/');
        });
      }
    });
  };

  componentWillUnmount() {
    Toast.hide();
  }

  render() {
    const {
      history: { goBack }
    } = this.props;
    const { hasMobileError, hasPwdError, mobile, pwd, loading } = this.state;
    return (
      <div>
        <NavBar
          mode="dark"
          leftContent={<Icon type="left" />}
          onLeftClick={goBack}
        >
          创建员工
        </NavBar>
        <List renderHeader={() => '输出员工手机号与密码'}>
          <InputItem
            type="phone"
            placeholder="请输入11位员工手机号"
            error={hasMobileError}
            onErrorClick={this.onMobileErrorClick}
            onChange={this.onMobileChange}
            value={mobile}
          >
            手机号
          </InputItem>
          <InputItem
            type="password"
            placeholder="请输入密码"
            error={hasPwdError}
            onErrorClick={this.onPwdErrorClick}
            onChange={this.onPwdChange}
            value={pwd}
          >
            密码
          </InputItem>
        </List>
        <WingBlank>
          <WhiteSpace />
          <Button
            type="primary"
            loading={loading}
            onClick={evt => this.onButtonClick(evt)}
          >
            创建
          </Button>
        </WingBlank>
      </div>
    );
  }
}

export default Create;
