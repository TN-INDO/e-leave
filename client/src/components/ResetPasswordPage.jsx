import React, { Component } from "react";
import { bindActionCreators } from "redux";
import { connect } from "react-redux";
import { NavLink } from "react-router-dom";
import {
  handleEdit,
  resetPassword
} from "../store/Actions/resetPasswordAction";
import { Layout, Form, Icon, Input, Button } from "antd";
const { Header, Content } = Layout;
const { Footer } = Layout;
const FormItem = Form.Item;

class ResetPasswordPage extends Component {
  constructor(props) {
    super(props);
    this.state = {
      loading: false,
      iconLoading: false
    };
  }

  componentWillMount() {
    console.log(" ----------------- Reset-Password ----------------- ");
  }

  componentDidMount() {
    if (
      localStorage.getItem("token") &&
      localStorage.getItem("role") === "admin"
    ) {
      this.props.history.push("/admin");
    } else if (
      localStorage.getItem("token") &&
      localStorage.getItem("role") === "director"
    ) {
      this.props.history.push("/director");
    } else if (
      localStorage.getItem("token") &&
      localStorage.getItem("role") === "supervisor"
    ) {
      this.props.history.push("/supervisor");
    } else if (
      localStorage.getItem("token") &&
      localStorage.getItem("role") === "employee"
    ) {
      this.props.history.push("/employee");
    }
    this.props.form.validateFields();
  }

  handleSumbitLogin = e => {
    e.preventDefault();

    this.props.form.validateFields((err, values) => {
      if (!err) {
        console.log("Received values of form: ", values);
      }
    });

    this.props.resetPassword(this.props.loginForm, url => {
      this.props.history.push(url);
    });
  };

  handleOnChangeLogin = e => {
    let loginForm = {
      ...this.props.loginForm,
      [e.target.name]: e.target.value
    };
    this.props.handleEdit(loginForm);

    this.props.form.setFieldsValue({
      [e.target.name]: e.target.value
    });
  };

  hasErrors(fieldsError) {
    return Object.keys(fieldsError).some(field => fieldsError[field]);
  }

  render() {
    const {
      getFieldDecorator,
      getFieldsError,
      getFieldError,
      isFieldTouched
    } = this.props.form;
    const emailError = isFieldTouched("email") && getFieldError("email");

    return (
      <div>
        <Layout className="App">
          <Header className="App-header">
            <div className="header-content">
              <h1>
                <Icon type="home" /> E-Leave
              </h1>
            </div>
          </Header>

          <Content
            className="container"
            style={{
              display: "flex",
              margin: "192px 15px ",
              justifyContent: "center"
            }}
          >
            <div
              style={{ padding: 50, background: "#fff", "border-radius": 7 }}
            >
              <Form onSubmit={this.handleSumbitLogin} className="login-form">
                <h1 align="left"> Forget Password ?</h1>
                <p align="left">
                  Enter your e-mail address below <br /> so we can send back
                  your password.
                </p>

                <FormItem
                  validateStatus={emailError ? "error" : ""}
                  help={emailError || ""}
                >
                  {getFieldDecorator("email", {
                    rules: [
                      {
                        type: "email",
                        message: "The input is not valid E-mail!"
                      },
                      {
                        required: true
                      }
                    ]
                  })(
                    <Input
                      type="email"
                      id="email"
                      name="email"
                      placeholder="email"
                      value={this.props.loginForm.email}
                      onChange={this.handleOnChangeLogin}
                      prefix={
                        <Icon
                          type="mail"
                          style={{ color: "rgba(0,0,0,.25)" }}
                        />
                      }
                    />
                  )}
                </FormItem>

                <FormItem>
                  <Button
                    type="primary"
                    htmlType="submit"
                    className="login-form-button"
                    onClick={this.handleSumbitLogin}
                    disabled={this.hasErrors(getFieldsError())}
                  >
                    FORGOT
                  </Button>
                </FormItem>
                <NavLink to="/">back!</NavLink>

                <FormItem />
              </Form>
            </div>
          </Content>

          <Footer className="Login-footer">
            <p>
              <a href="http://www.tnis.com">PT. TNIS Service Indonesia</a>{" "}
              &copy; 2018. All Right Reserved.
            </p>
          </Footer>
        </Layout>
      </div>
    );
  }
}

const mapStateToProps = state => ({
  loginForm: state.resetPasswordReducer
});

const mapDispatchToProps = dispatch =>
  bindActionCreators(
    {
      handleEdit,
      resetPassword
    },
    dispatch
  );

const WrappedLoginForm = Form.create()(ResetPasswordPage);

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(WrappedLoginForm);
