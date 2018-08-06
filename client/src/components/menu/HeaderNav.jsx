import React from "react";
import { NavLink, withRouter } from "react-router-dom";
import { Layout, Menu, Button, Icon } from "antd";
const { Header } = Layout;
const SubMenu = Menu.SubMenu;

const ButtonLogout = withRouter(({ history }) => (
  <Button
    onClick={() => {
      localStorage.clear();
      history.push("/");
    }}
    type="danger"
    ghost
  >
    Logout
  </Button>
));

export default class HeaderNav extends React.Component {
  constructor(props) {
    super(props);
    this.state = { role: localStorage.getItem("role") };
  }

  render() {
    if (this.state.role === "employee") {
      return (
        <Header>
          <Menu
            theme="dark"
            mode="horizontal"
            defaultSelectedKeys={["1"]}
            style={{ lineHeight: "64px" }}
          >
            <Menu.Item key="1">
              <NavLink to="/employee">
                <span>
                  <Icon type="home" />
                </span>
                Home
              </NavLink>
            </Menu.Item>

            <Menu.Item key="2">
              <NavLink to="/profile">
                <span>
                  <Icon type="profile" />
                </span>
                My Profile
              </NavLink>
            </Menu.Item>

            <SubMenu
              title={
                <span>
                  <Icon type="user" />Employee
                </span>
              }
            >
              <Menu.Item key="employee:1">
                <NavLink to="/employee/leave-request">
                  <span>
                    <Icon type="form" />
                  </span>
                  Form Request Leave
                </NavLink>
              </Menu.Item>
              <Menu.Item key="employee:2">
                <NavLink to="/employee/list-pending-request">
                  <span>
                    <Icon type="schedule" />
                  </span>List Pending Request
                </NavLink>
              </Menu.Item>
              <Menu.Item key="employee:3">
                <NavLink to="/employee/list-approve-request">
                  <span>
                    <Icon type="schedule" />
                  </span>List Approve Request
                </NavLink>
              </Menu.Item>
              <Menu.Item key="employee:4">
                <NavLink to="/employee/list-reject-request">
                  <span>
                    <Icon type="schedule" />
                  </span>List Reject Request
                </NavLink>
              </Menu.Item>
            </SubMenu>

            <Menu.Item key="4">
              <ButtonLogout />
            </Menu.Item>
          </Menu>
          <div
            style={{
              display: "flex",
              justifyContent: "space-between"
            }}
          />
        </Header>
      );
    } else if (this.state.role === "supervisor") {
      return (
        <Header>
          <Menu
            theme="dark"
            mode="horizontal"
            defaultSelectedKeys={["1"]}
            style={{ lineHeight: "64px" }}
          >
            <Menu.Item key="1">
              <NavLink to="/">
                <span>
                  <Icon type="home" />
                </span>
                Home
              </NavLink>
            </Menu.Item>

            <Menu.Item key="2">
              <NavLink to="/profile">
                <span>
                  <Icon type="profile" />
                </span>
                My Profile
              </NavLink>
            </Menu.Item>

            <SubMenu
              title={
                <span>
                  <Icon type="user" />Supervisor
                </span>
              }
            >
              <Menu.Item key="supervisor:1">
                <NavLink to="/supervisor/list-pending-request">
                  <span>
                    <Icon type="schedule" />
                  </span>List Pending Request
                </NavLink>
              </Menu.Item>
              <Menu.Item key="supervisor:2">
                <NavLink to="/supervisor/list-approve-request">
                  <span>
                    <Icon type="schedule" />
                  </span>List Approve Request
                </NavLink>
              </Menu.Item>
              <Menu.Item key="supervisor:3">
                <NavLink to="/supervisor/list-reject-request">
                  <span>
                    <Icon type="schedule" />
                  </span>List Reject Request
                </NavLink>
              </Menu.Item>
            </SubMenu>

            <SubMenu
              title={
                <span>
                  <Icon type="user" />Employee
                </span>
              }
            >
              <Menu.Item key="employee:1">
                <NavLink to="/employee/leave-request">
                  <span>
                    <Icon type="form" />
                  </span>
                  Form Request Leave
                </NavLink>
              </Menu.Item>
              <Menu.Item key="employee:2">
                <NavLink to="/employee/list-pending-request">
                  <span>
                    <Icon type="schedule" />
                  </span>List Pending Request
                </NavLink>
              </Menu.Item>
              <Menu.Item key="employee:3">
                <NavLink to="/employee/list-approve-request">
                  <span>
                    <Icon type="schedule" />
                  </span>List Approve Request
                </NavLink>
              </Menu.Item>
              <Menu.Item key="employee:4">
                <NavLink to="/employee/list-reject-request">
                  <span>
                    <Icon type="schedule" />
                  </span>List Reject Request
                </NavLink>
              </Menu.Item>
            </SubMenu>

            <Menu.Item key="10">
              <ButtonLogout />
            </Menu.Item>
          </Menu>
          <div
            style={{
              display: "flex",
              justifyContent: "space-between"
            }}
          />
        </Header>
      );
    } else if (this.state.role === "director") {
      return (
        <Header>
          <Menu
            theme="dark"
            mode="horizontal"
            defaultSelectedKeys={["1"]}
            style={{ lineHeight: "64px" }}
          >
            <Menu.Item key="1">
              <NavLink to="/">
                <span>
                  <Icon type="home" />
                </span>
                Home
              </NavLink>
            </Menu.Item>

            <Menu.Item key="2">
              <NavLink to="/profile">
                <span>
                  <Icon type="profile" />
                </span>
                My Profile
              </NavLink>
            </Menu.Item>

            <SubMenu
              title={
                <span>
                  <Icon type="user" />Director
                </span>
              }
            >
              <Menu.Item key="user:1">
                <NavLink to="/director/list-pending-request">
                  <span>
                    <Icon type="schedule" />
                  </span>List Pending Request
                </NavLink>
              </Menu.Item>
              <Menu.Item key="user:2">
                <NavLink to="/director/list-approve-request">
                  <span>
                    <Icon type="schedule" />
                  </span>List Approve Request
                </NavLink>
              </Menu.Item>
              <Menu.Item key="user:3">
                <NavLink to="/director/list-reject-request">
                  <span>
                    <Icon type="schedule" />
                  </span>List Reject Request
                </NavLink>
              </Menu.Item>
            </SubMenu>

            <Menu.Item key="3">
              <ButtonLogout />
            </Menu.Item>
          </Menu>
          <div
            style={{
              display: "flex",
              justifyContent: "space-between"
            }}
          />
        </Header>
      );
    } else {
      return null;
    }
  }
}
