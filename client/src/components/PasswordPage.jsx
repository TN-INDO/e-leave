import React, { Component } from "react";
import { bindActionCreators } from "redux";
import { connect } from "react-redux";
import { handleEdit, updateNewPassword } from "../store/Actions/passwordAction";
import HeaderNav from "./menu/HeaderNav";
import Footer from "./menu/Footer";
import { Layout, Form, Input, Button } from "antd";
const { Content } = Layout;
const FormItem = Form.Item;

class PasswordPage extends Component {
  constructor(props) {
    super(props);
    this.state = {
      from: null,
      to: null,
      endOpen: false,
      confirmDirty: false
    };

    this.handleOnChange = this.handleOnChange.bind(this);
  }

  componentWillMount() {
    console.log(" ----------------- Update-Password ----------------- ");
    if (!localStorage.getItem("token")) {
      this.props.history.push("/");
    } else if (
      localStorage.getItem("role") !== "employee" &&
      localStorage.getItem("role") !== "supervisor" &&
      localStorage.getItem("role") !== "director" &&
      localStorage.getItem("role") !== "admin"
    ) {
      this.props.history.push("/");
    }
  }

  updatePassword = e => {
    e.preventDefault();
    this.props.updateNewPassword(this.props.passwordForm, url => {
      this.props.history.push(url);
    });
    this.props.form.validateFieldsAndScroll((err, values) => {
      if (!err) {
        console.log("Received values of form: ", values);
      }
    });
  };

  handleOnChange = e => {
    let newPassword = {
      ...this.props.passwordForm,
      [e.target.name]: e.target.value
    };
    this.props.handleEdit(newPassword);
  };

  handleConfirmBlur = e => {
    const value = e.target.value;
    this.setState({ confirmDirty: this.state.confirmDirty || !!value });
  };

  compareToFirstPassword = (rule, value, callback) => {
    const form = this.props.form;
    if (value && value !== form.getFieldValue("password")) {
      callback("Wrong confirm password!");
    } else {
      callback();
    }
  };

  validateToNextPassword = (rule, value, callback) => {
    const form = this.props.form;
    if (value && this.state.confirmDirty) {
      form.validateFields(["confirm"], { force: true });
    }
    callback();
  };

  render() {
    const { getFieldDecorator } = this.props.form;
    const formItemLayout = {
      labelCol: {
        xs: { span: 24 },
        sm: { span: 10 }
      },
      wrapperCol: {
        xs: { span: 24 },
        sm: { span: 10 }
      },
      style: {}
    };

    const formStyle = {
      width: "100%"
    };

    return (
      <Layout>
        <HeaderNav />
        <Content
          className="container"
          style={{
            display: "flex",
            margin: "20px 16px 0",
            justifyContent: "center",
            paddingBottom: "411px"
          }}
        >
          <div
            style={{
              padding: 50,
              paddingBottom: 50,
              paddingTop: 50,
              background: "#fff"
              // minHeight: 360
            }}
          >
            <h1> UPDATE PASSWORD </h1>

            <Form className="login-form">
              <FormItem {...formItemLayout} label="Old Password">
                {getFieldDecorator("old password", {
                  rules: [
                    {
                      required: true
                    },
                    { min: 7, message: "password length minimum is 7" }
                  ]
                })(
                  <Input
                    type="password"
                    id="old_password"
                    name="old_password"
                    placeholder="old password"
                    onChange={this.handleOnChange}
                    style={formStyle}
                  />
                )}
              </FormItem>

              <FormItem {...formItemLayout} label="New Password">
                {getFieldDecorator("password", {
                  rules: [
                    {
                      required: true
                    },
                    { min: 7, message: "password length minimum is 7" },
                    {
                      validator: this.validateToNextPassword
                    }
                  ]
                })(
                  <Input
                    type="password"
                    id="new_password"
                    name="new_password"
                    placeholder="new password"
                    onChange={this.handleOnChange}
                    style={formStyle}
                  />
                )}
              </FormItem>

              <FormItem {...formItemLayout} label="Confirm Password">
                {getFieldDecorator("confirm", {
                  rules: [
                    {
                      required: true
                    },
                    { min: 7, message: "password length minimum is 7" },
                    {
                      validator: this.compareToFirstPassword
                    }
                  ]
                })(
                  <Input
                    type="password"
                    id="confirm_password"
                    name="confirm_password"
                    placeholder="confirm password"
                    onChange={this.handleOnChange}
                    style={formStyle}
                  />
                )}
              </FormItem>

              <FormItem>
                <Button
                  onClick={this.updatePassword}
                  htmlType="submit"
                  type="primary"
                  style={{
                    width: "35%"
                  }}
                >
                  Update
                </Button>
              </FormItem>
            </Form>
          </div>
        </Content>

        <Footer />
      </Layout>
    );
  }
}

const mapStateToProps = state => ({
  passwordForm: state.passwordReducer
});

const WrappedPasswordForm = Form.create()(PasswordPage);

const mapDispatchToProps = dispatch =>
  bindActionCreators(
    {
      handleEdit,
      updateNewPassword
    },
    dispatch
  );

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(WrappedPasswordForm);
