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

export default class HeaderAdmin extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      role: localStorage.getItem("role"),
      id: localStorage.getItem("id")
    };
  }
  render() {
    return (
      <Header>
        <Menu
          theme="dark"
          // mode={mobileVersion ? 'vertical' : 'horizontal'}
          mode="horizontal"
          defaultSelectedKeys={["1"]}
          style={{ lineHeight: "64px" }}
        >
          <Menu.Item key="1">
            <NavLink to="/admin">
              <span>
                <Icon type="home" />
              </span>Home
            </NavLink>
          </Menu.Item>

          <Menu.Item key="3">
            <NavLink to="/admin/register-user">
              <span>
                <Icon type="form" />
              </span>Add User
            </NavLink>
          </Menu.Item>
          <SubMenu
            title={
              <span>
                <Icon type="user" />Employee
              </span>
            }
          >
            <Menu.Item key="schedule:1">
              <NavLink to="/admin/list-pending-request">
                <span>
                  <Icon type="schedule" />
                </span>List Pending Request
              </NavLink>
            </Menu.Item>
            <Menu.Item key="schedule:2">
              <NavLink to="/admin/list-approve-request">
                <span>
                  <Icon type="schedule" />
                </span>List Approve Request
              </NavLink>
            </Menu.Item>
            <Menu.Item key="schedule:3">
              <NavLink to="/admin/list-reject-request">
                <span>
                  <Icon type="schedule" />
                </span>List Reject Request
              </NavLink>
            </Menu.Item>
          </SubMenu>

          <SubMenu
            title={
              <span>
                <Icon type="setting" />Settings
              </span>
            }
          >
            {/* <Menu.Item key="setting:1">
              <NavLink to="/admin/profile">
                <span>
                  <Icon type="profile" />
                </span>Profile
              </NavLink>
            </Menu.Item> */}
            <Menu.Item key="setting:2">
              <NavLink to="/admin/edit-password">
                <span>
                  <Icon type="form" />
                </span>Change Password
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
  }
}
