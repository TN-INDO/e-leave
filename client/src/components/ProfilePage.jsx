import React, { Component } from "react";
import { connect } from "react-redux";
import { bindActionCreators } from "redux";
import { profileFetchData } from "../store/Actions/profileAction";
import { getSupervisors } from "../store/Actions/AddSupervisorAction";
import moment from "moment-timezone";
import { Layout, Button, Form, Input, Select, DatePicker } from "antd";
import HeaderNav from "./menu/HeaderNav";
import Footer from "./menu/Footer";
const { Content } = Layout;
const FormItem = Form.Item;
const Option = Select.Option;

class ProfilePage extends Component {
  componentWillMount() {
    console.log(" ----------------- Profile-Page ----------------- ");
  }

  componentDidMount() {
    if (!localStorage.getItem("token")) {
      this.props.history.push("/");
    } else if (
      localStorage.getItem("role") !== "employee" &&
      localStorage.getItem("role") !== "supervisor" &&
      localStorage.getItem("role") !== "director"
    ) {
      this.props.history.push("/");
      window.location.href = "/";
    }

    this.props.profileFetchData();
    this.props.getSupervisors();
  }

  editPassword = (user, employeeNumber) => {
    this.props.history.push({
      pathname: "/profile/" + employeeNumber,
      state: { user: user }
    });
  };

  handleBlur() {
    console.log("blur");
  }

  handleFocus() {
    console.log("focus");
  }

  render() {
    const formItemLayout = {
      labelCol: {
        xs: { span: 24 },
        sm: { span: 8 }
      },
      wrapperCol: {
        xs: { span: 24 },
        sm: { span: 16 }
      }
    };

    const formStyle = {
      width: "100%"
    };

    const dateFormat = "DD-MM-YYYY";
    let supervisorName;
    if (this.props.supervisor) {
      this.props.supervisor.map(d => {
        if (d.supervisor_id === this.props.user.supervisor_id) {
          supervisorName = d.name;
        }
        return d;
      });
    }

    return (
      <div>
        <Layout>
          <HeaderNav />
          <Content
            className="container"
            style={{
              display: "flex",
              margin: "20px 16px 0",
              justifyContent: "space-around",
              paddingBottom: "179px"
            }}
          >
            <div
              style={{
                padding: 100,
                paddingBottom: 50,
                paddingTop: 50,
                background: "#fff"
              }}
            >
              <h1> My Profile </h1>
              <div>
                <Form>
                  <FormItem {...formItemLayout} label="Name">
                    <Input
                      type="text"
                      id="name"
                      name="name"
                      placeholder="name"
                      value={this.props.user.name}
                    />
                  </FormItem>

                  <FormItem {...formItemLayout} label="Email">
                    <Input
                      type="email"
                      id="email"
                      name="email"
                      value={this.props.user.email}
                    />
                  </FormItem>

                  <FormItem {...formItemLayout} label="Gender">
                    <Select
                      id="gender"
                      name="gender"
                      optionFilterProp="children"
                      value={this.props.user.gender}
                      onFocus={this.handleFocus}
                      onBlur={this.handleBlur}
                    >
                      <Option value="male" disabled>
                        Male
                      </Option>
                      <Option value="female" disabled>
                        Female
                      </Option>
                    </Select>
                  </FormItem>

                  <FormItem {...formItemLayout} label="Position">
                    <Input
                      type="text"
                      id="position"
                      name="position"
                      value={this.props.user.position}
                    />
                  </FormItem>

                  <FormItem {...formItemLayout} label="Start Working Date">
                    <DatePicker
                      id="start_working_date"
                      name="start_working_date"
                      format={dateFormat}
                      value={moment(
                        this.props.user.start_working_date,
                        dateFormat
                      )}
                      style={formStyle}
                      disabled
                    />
                  </FormItem>

                  <FormItem {...formItemLayout} label="Mobile Phone">
                    <Input
                      type="text"
                      id="mobile_phone"
                      name="mobile_phone"
                      value={this.props.user.mobile_phone}
                    />
                  </FormItem>

                  <FormItem {...formItemLayout} label="Role">
                    <Select
                      id="role"
                      name="role"
                      optionFilterProp="children"
                      onFocus={this.handleFocus}
                      onBlur={this.handleBlur}
                      value={this.props.user.role}
                    >
                      <Option value="employee" disabled>
                        Employee
                      </Option>
                      <Option value="supervisor" disabled>
                        Supervisor
                      </Option>
                      <Option value="dirctor" disabled>
                        Director
                      </Option>
                    </Select>
                  </FormItem>

                  {this.props.user.role === "employee" ? (
                    <FormItem {...formItemLayout} label="Supervisor">
                      <Select
                        id="supervisor_id"
                        name="supervisor_id"
                        optionFilterProp="children"
                        onFocus={this.handleFocus}
                        onBlur={this.handleBlur}
                        value={supervisorName}
                      >
                        {this.props.supervisor.map(d => (
                          <Option key={d.supervisor_id} disabled>
                            {d.name}
                          </Option>
                        ))}
                      </Select>
                    </FormItem>
                  ) : (
                    ""
                  )}

                  {this.props.user.updated_at !== "0001-01-01T00:00:00Z" ? (
                    <FormItem {...formItemLayout} label="Updated At">
                      <Input
                        type="text"
                        id="updated_at"
                        name="updated_at"
                        value={moment(this.props.user.updated_at)
                          .tz("Asia/Jakarta")
                          .format("DD-MM-YYYY HH:mm")}
                      />
                    </FormItem>
                  ) : (
                    ""
                  )}

                  <FormItem>
                    <Button
                      onClick={() => {
                        this.editPassword(
                          this.props.user,
                          this.props.user.employee_number
                        );
                      }}
                      htmlType="edit"
                      type="primary"
                    >
                      Edit Password
                    </Button>
                  </FormItem>
                </Form>
              </div>
            </div>
          </Content>

          <Footer />
        </Layout>
      </div>
    );
  }
}

const mapStateToProps = state => ({
  loading: state.profileReducer.loading,
  user: state.profileReducer.user,
  supervisor: state.AddSupervisorReducer.supervisor
});

const WrappedPofileForm = Form.create()(ProfilePage);

const mapDispatchToProps = dispatch =>
  bindActionCreators(
    {
      profileFetchData,
      getSupervisors
    },
    dispatch
  );

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(WrappedPofileForm);
