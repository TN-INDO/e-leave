import React, { Component } from 'react';
import { BrowserRouter, Route, Switch } from 'react-router-dom';
import { Provider } from 'react-redux';
import store from './store';
import './App.css';

import LoginPage from './components/LoginPage';
import ResetPasswordPage from './components/ResetPasswordPage';

import AdminLandingPage from './components/AdminLandingPage';
import RegisterPage from './components/RegisterPage';
import AdminEditPage from './components/AdminEditPage';
import AdminEditBalancePage from './components/AdminEditBalancePage';
import AdminPendingPage from './components/AdminPendingPage';
import AdminApprovePage from './components/AdminApprovePage';
import AdminRejectPage from './components/AdminRejectPage';

import EmployeeLandingPage from './components/EmployeeLandingPage';
import ProfilePage from './components/ProfilePage';
import PasswordPage from './components/PasswordPage';
import PasswordPageAdmin from './components/PasswordPageAdmin';

import LeaveRequestPage from './components/LeaveRequestPage';
import LeaveEditPage from './components/LeaveEditPage';
import EmployeePendingPage from './components/EmployeePendingPage';
import EmployeeApprovePage from './components/EmployeeApprovePage';
import EmployeeRejectPage from './components/EmployeeRejectPage';

import SupervisorLandingPage from './components/SupervisorLandingPage';
import SupervisorPendingPage from './components/SupervisorPendingPage';
import SupervisorApprovePage from './components/SupervisorApprovePage';
import SupervisorRejectPage from './components/SupervisorRejectPage';

import DirectorLandingPage from './components/DirectorLandingPage';
import DirectorPendingPage from './components/DirectorPendingPage';
import DirectorApprovePage from './components/DirectorApprovePage';
import DirectorRejectPage from './components/DirectorRejectPage';

import Notfound from './components/NotFound';

class App extends Component {
  render() {
    return (
      <Provider store={store}>
        <div className="App">
          <BrowserRouter>
            <Switch>
              <Route exact path="/" component={LoginPage} />
              <Route exact path="/login" component={LoginPage} />
              <Route exact path="/reset-password" component={ResetPasswordPage} />

              <Route exact path="/admin" component={AdminLandingPage} />
              <Route exact path="/admin/register-user" component={RegisterPage} />            
              <Route exact path={`/admin/edit-user/:id`} component={AdminEditPage}/>
              <Route exact path={`/admin/edit-balance/:id`} component={AdminEditBalancePage}/>
              <Route path={`/admin/edit-password`} component={PasswordPageAdmin}/>
              <Route exact path="/admin/list-pending-request" component={AdminPendingPage} />
              <Route exact path="/admin/list-approve-request" component={AdminApprovePage} />              
              <Route exact path="/admin/list-reject-request" component={AdminRejectPage} />              
                           
              <Route exact path="/employee" component={EmployeeLandingPage} />

              <Route exact path="/profile" component={ProfilePage} />
              <Route path={`/profile/:id`} component={PasswordPage}/>

              <Route exact path="/employee/leave-request" component={LeaveRequestPage} />                          
              <Route exact path={`/editrequest/:id`} component={LeaveEditPage}/>
              <Route exact path="/employee/list-pending-request" component={EmployeePendingPage} />
              <Route exact path="/employee/list-approve-request" component={EmployeeApprovePage} />
              <Route exact path="/employee/list-reject-request" component={EmployeeRejectPage} />

              <Route exact path="/supervisor" component={SupervisorLandingPage} />
              <Route exact path="/supervisor/list-pending-request" component={SupervisorPendingPage} />
              <Route exact path="/supervisor/list-approve-request" component={SupervisorApprovePage} />
              <Route exact path="/supervisor/list-reject-request" component={SupervisorRejectPage} />

              <Route exact path="/director" component={DirectorLandingPage} />
              <Route exact path="/director/list-pending-request" component={DirectorPendingPage} />
              <Route exact path="/director/list-approve-request" component={DirectorApprovePage} />
              <Route exact path="/director/list-reject-request" component={DirectorRejectPage} />
              
              <Route component={Notfound} />
            </Switch>
          </BrowserRouter>
        </div>
      </Provider>
    );
  }
}
export default App;
