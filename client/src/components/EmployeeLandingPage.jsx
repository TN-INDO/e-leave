import React from "react";
import { bindActionCreators } from "redux";
import { connect } from "react-redux";
import { userSummaryFetchData } from "../store/Actions/userSummaryAction";
import HeaderNav from "./menu/HeaderNav";
import Loading from "./menu/Loading";
import Footer from "./menu/Footer";
import { Layout, Row, Col } from "antd";
const { Content } = Layout;

export class EmployeeLandingPage extends React.Component {
  componentWillMount() {
    console.log(" ----------------- Employee-Landing-Page ----------------- ");
  }

  componentDidMount() {
    if (!localStorage.getItem("token")) {
      this.props.history.push("/");
    } else if (localStorage.getItem("role") !== "employee") {
      this.props.history.push("/");
    }
    this.props.userSummaryFetchData();
  }

  render() {
    const summary = [];
    const typeLeave = [];
    const dataSummary = this.props.userSummary;
    const dataType = this.props.userType;

    if (dataSummary) {
      for (let i = 0; i < dataSummary.length; i++) {
        summary.push(
          <p>
            {dataSummary[i].type_name}: {dataSummary[i].used} day(s)
          </p>
        );
      }
    } else {
      summary.push(<p />);
    }

    for (let j in dataType) {
      if (dataType[j].type_name === "Annual Leave") {
        typeLeave.push(
          <p>
            {`${dataType[j].type_name} : ${
              dataType[j].leave_remaining
            } days of 12 days`}
          </p>
        );
      } else if (dataType[j].type_name === "Errand Leave") {
        typeLeave.push(
          <p>{`${dataType[j].type_name} : ${
            dataType[j].leave_remaining
          } days of 3 days`}</p>
        );
      } else if (dataType[j].type_name === "Sick Leave") {
        typeLeave.push(
          <p>{`${dataType[j].type_name} (up to 10 days) : ${
            dataType[j].leave_remaining
          } days of 30 days`}</p>
        );
      } else if (dataType[j].type_name === "Marriage Leave") {
        typeLeave.push(
          <p>
            {`${dataType[j].type_name} : ${
              dataType[j].leave_remaining
            } days of 2 days`}{" "}
          </p>
        );
      } else if (dataType[j].type_name === "Maternity Leave") {
        typeLeave.push(
          <p>
            {`${dataType[j].type_name} : ${
              dataType[j].leave_remaining
            } days of 90 days`}
          </p>
        );
      } else if (dataType[j].type_name === "Other Leave") {
        typeLeave.push(
          <p>
            {`${dataType[j].type_name} : ${
              dataType[j].leave_remaining
            } days of 2 days`}
          </p>
        );
      }
    }

    if (this.props.loading) {
      return <Loading />;
    } else {
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
                paddingBottom: "542px"
              }}
            >
              <div style={{ padding: 40, width: "50%", background: "#fff" }}>
                <Row>
                  <Col span={12}>
                    <div
                      style={{
                        float: "left"
                      }}
                    >
                      <h3>
                        {" "}
                        Used so far
                        <hr
                          className="id"
                          style={{
                            borderBottom: "1px solid black",
                            width: "276px"
                          }}
                        />
                      </h3>

                      {summary}
                    </div>
                  </Col>
                  <Col span={12}>
                    <div
                      style={{
                        float: "right"
                      }}
                    >
                      <h3>
                        Available types
                        <hr
                          className="id"
                          style={{
                            borderBottom: "1px solid black",
                            width: "276px"
                          }}
                        />
                      </h3>
                      {typeLeave}
                    </div>
                  </Col>
                </Row>
              </div>
            </Content>

            <Footer />
          </Layout>
        </div>
      );
    }
  }
}
const mapStateToProps = state => ({
  userSummary: state.fetchUserSummaryReducer.userSummary,
  loading: state.fetchUserSummaryReducer.loading,
  userType: state.fetchUserSummaryReducer.userType
});

const mapDispatchToProps = dispatch =>
  bindActionCreators(
    {
      userSummaryFetchData
    },
    dispatch
  );

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(EmployeeLandingPage);
