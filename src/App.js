import React, { PureComponent } from 'react';
import { Link } from 'react-router-dom';
import { NavBar, Toast, List } from 'antd-mobile';

import { employeeList } from './service/employee';

const Item = List.Item;
const Brief = Item.Brief;

class App extends PureComponent {
  constructor(props) {
    super(props);
    this.state = {
      count: 0,
      employeeList: []
    };
  }
  componentDidMount() {
    this.loadData();
  }

  loadData = () => {
    Toast.loading('数据加载中...', 0);
    employeeList().then(res => {
      const { data } = res;
      if (data.code !== "0") {
        Toast.fail(data.message);
      } else {
        Toast.hide();
        this.setState({
          count: data.count,
          employeeList: data.employeeList
        });
      }
    });
  };

  handleRefresh = evt => {
    if (evt !== undefined && evt.preventDefault) evt.preventDefault();
    this.loadData();
  };

  render() {
    const { employeeList } = this.state;
    const {
      history: { push }
    } = this.props;
    return (
      <div className="App">
        <NavBar
          mode="dark"
          leftContent="刷新"
          rightContent={
            <Link to="/create" style={{ color: 'white' }}>
              创建
            </Link>
          }
          onLeftClick={evt => this.handleRefresh(evt)}
        >
          YQ签到神器
        </NavBar>
        {!employeeList.length ? (
          <List>
            <Item arrow="horizontal" onClick={() => push('/create')}>
              请创建员工信息
            </Item>
          </List>
        ) : (
          <List>
            {employeeList.map(employee => (
              <Item
                key={employee.ID}
                arrow="horizontal"
                multipleLine
                onClick={() => push(`/modify/${employee.ID}`)}
              >
                {employee.Name}
                <Brief>
                  员工ID：
                  {employee.EmployeeID}
                </Brief>
              </Item>
            ))}
          </List>
        )}
      </div>
    );
  }
}

export default App;
