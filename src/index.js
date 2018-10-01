import React from 'react';
import ReactDOM from 'react-dom';
import {
  BrowserRouter as Router,
  Switch,
  Route,
  Redirect
} from 'react-router-dom';
import 'antd-mobile/dist/antd-mobile.less';
import './index.less';
import App from './App';
import Create from './Create';
import Modify from './Modify';
import registerServiceWorker from './registerServiceWorker';

ReactDOM.render(
  <Router>
    <Switch>
      <Route path="/" exact component={App} />
      <Route path="/create" exact component={Create} />
      <Route path="/modify/:id" exact component={Modify} />
      <Redirect to="/" />
    </Switch>
  </Router>,
  document.getElementById('root')
);
registerServiceWorker();
