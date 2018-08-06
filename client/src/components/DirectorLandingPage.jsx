import React, { Component } from "react";
import { bindActionCreators } from "redux";
import { connect } from "react-redux";
import { fetchAdminLeaveApprove } from "../store/Actions/adminActions";
import FullCalendar from "fullcalendar-reactwrapper";
import moment from "moment";
import "fullcalendar/dist/fullcalendar.css";
import HeaderNav from "./menu/HeaderNav";
import Loading from "./menu/Loading";
import Footer from "./menu/Footer";
import { Layout, message } from "antd";
const { Content } = Layout;

export class DirectorLandingPage extends Component {
  constructor(props) {
    super(props);
    this.state = {
      leaves: this.props.leave
    };
  }

  componentWillMount() {
    console.log(" ----------------- Director-Landing-Page ----------------- ");
  }

  componentDidMount() {
    if (!localStorage.getItem("token")) {
      this.props.history.push("/");
    } else if (localStorage.getItem("role") !== "director") {
      this.props.history.push("/");
    }
    this.props.fetchAdminLeaveApprove();
  }

  componentWillReceiveProps(nextProps) {
    if (nextProps.leave !== this.props.leave) {
      this.setState({
        leaves: nextProps.leave
      });
    }
  }

  formatDate(value) {
    const date = new Date(value._d),
      mnth = ("0" + (date.getMonth() + 1)).slice(-2),
      day = ("0" + date.getDate()).slice(-2);
    let dateFrom = [date.getFullYear(), mnth, day].join("-");
    return dateFrom;
  }

  eventRender(event, element) {
    if (event.className === "Sick Leave") {
      element.css({
        "background-color": "#333333",
        "border-color": "#333333"
      });
    }
  }

  render() {
    const events = [];
    if (this.state.leaves) {
      for (let i = 0; i < this.state.leaves.length; i++) {
        events.push({
          title: `${this.state.leaves[i].name} ${this.state.leaves[i].total} day(s) for ${
            this.state.leaves[i].type_name
          }`,
          start: this.formatDate(
            moment(this.state.leaves[i].date_from, "DD-MM-YYYY")
          ),
          end: this.formatDate(
            moment(this.state.leaves[i].date_to, "DD-MM-YYYY").add(1, "d")
          ),
          className: this.state.leaves[i].type_name
        });
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
                paddingBottom: "20px"
              }}
            >
              <div style={{ padding: 20, background: "#fff"}}>
                <FullCalendar
                  id="your-custom-ID"
                  header={{
                    left: "prev,next today myCustomButton",
                    center: "title",
                    right: "month,basicWeek,basicDay"
                  }}
                  navLinks={true} 
                  editable={true}
                  eventLimit={true}
                  events={events}
                  eventColor={"#378006"}
                  eventClick={function(calEvent, jsEvent, view, resourceObj) {
                    message.info(calEvent.title);
                  }}
                />
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
  loading: state.adminReducer.loading,
  leave: state.adminReducer.leaves
});

const mapDispatchToProps = dispatch =>
  bindActionCreators(
    {
      fetchAdminLeaveApprove
    },
    dispatch
  );

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(DirectorLandingPage);
