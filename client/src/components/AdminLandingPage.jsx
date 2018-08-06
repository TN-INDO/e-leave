import React, { Component } from "react";
import { bindActionCreators } from "redux";
import { connect } from "react-redux";
import { adminGetUsers, adminDeleteUser } from "../store/Actions/adminActions";
import HeaderNav from "./menu/HeaderAdmin";
import Loading from "./menu/Loading";
import Footer from "./menu/Footer";
import { Layout, Table, Button, Divider, Popconfirm, message } from "antd";
const { Content } = Layout;

class AdminLandingPage extends Component {
  constructor(props) {
    super(props);
    this.columns = [
      {
        title: "Employee Number",
        dataIndex: "employee_number",
        key: "employee_number",
        width: "10%"
      },
      {
        title: "Name",
        dataIndex: "name",
        key: "name",
        width: "18%"
      },
      {
        title: "Email",
        dataIndex: "email",
        key: "email",
        width: "18%"
      },
      {
        title: "Position",
        dataIndex: "position",
        key: "position",
        width: "15%"
      },
      {
        title: "Role",
        dataIndex: "role",
        key: "role",
        width: "10%"
      },
      {
        title: "Action",
        key: "action",
        width: "25%",
        render: (text, record) => (
          //  (record.employee_number))
          <span>
            <Button
              onClick={() => {
                this.editUser(this.props.users, record.employee_number);
              }}
              type="primary"
            >
              Edit
            </Button>

            <Divider type="vertical" />
            <Popconfirm
              placement="top"
              title={"Are you sure delete this employee?"}
              onConfirm={() => {
                this.adminDeleteUser(this.props.users, record.employee_number);
                message.success("Employee has been delete!");
              }}
              okText="Yes"
              cancelText="No"
            >
              <Button type="danger">Delete</Button>
            </Popconfirm>
            <Divider type="vertical" />
            {/* {record.role === "supervisor" || record.role === "employee" ? (
              <Button
                onClick={() => {
                  this.editBalance(record.employee_number);
                }}
                type="primary"
              >
                Edit Leave Balance
              </Button>
            ) : (
              ""
            )} */}
          </span>
        )
      }
    ];
  }

  componentWillMount() {
    console.log(" ----------------- Admin-Page ----------------- ");
  }

  componentDidMount() {
    if (localStorage.getItem("role") !== "admin") {
      this.props.history.push("/");
    }
    this.props.adminGetUsers();
  }

  editUser = (users, employeeNumber) => {
    this.props.history.push({
      pathname: "/admin/edit-user/" + employeeNumber,
      state: { users: users }
    });
  };

  editBalance = employeeNumber => {
    this.props.history.push({
      pathname: "/admin/edit-balance/" + employeeNumber
    });
  };

  adminDeleteUser = (users, employeeNumber) => {
    this.props.adminDeleteUser(users, employeeNumber);
  };

  onShowSizeChange(current, pageSize) {
    console.log(current, pageSize);
  }

  render() {
    if (this.props.loading) {
      return <Loading />;
    } else {
      return (
        <Layout>
          <HeaderNav />

          <Content
            className="container"
            style={{
              display: "flex",
              margin: "20px 16px 0",
              justifyContent: "center",
              paddingBottom: "627px"
            }}
          >
            <div style={{ padding: 20, background: "#fff" }}>
              <Table
                columns={this.columns}
                dataSource={this.props.users}
                rowKey={record => record.employee_number}
                pagination={{
                  className: "my-pagination",
                  defaultCurrent: 1,
                  defaultPageSize: 5,
                  total: 50,
                  showSizeChanger: this.onShowSizeChange
                }}
              />
            </div>
          </Content>
          <Footer />
        </Layout>
      );
    }
  }
}

const mapStateToProps = state => ({
  loading: state.adminReducer.loading,
  users: state.adminReducer.users
});

const mapDispatchToProps = dispatch =>
  bindActionCreators(
    {
      adminGetUsers,
      adminDeleteUser
    },
    dispatch
  );

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(AdminLandingPage);
