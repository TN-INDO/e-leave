import React, { Component } from "react";
import { connect } from "react-redux";
import { bindActionCreators } from "redux";
import update from "react-addons-update";
import moment from "moment-business-days";
import {
  fetchedEdit,
  handleEdit,
  saveEditLeave
} from "../store/Actions/editRequestAction";
import { typeLeaveFetchData } from "../store/Actions/typeLeaveAction";
import {
  Layout,
  Button,
  Form,
  Input,
  Select,
  Checkbox,
  DatePicker
} from "antd";
import HeaderNav from "./menu/HeaderNav";
import Loading from "./menu/Loading";
import Footer from "./menu/Footer";
const { Content } = Layout;
const { TextArea } = Input;
const FormItem = Form.Item;
const Option = Select.Option;

class LeaveEditPage extends Component {
  constructor(props) {
    super(props);
    this.state = {
      from: null,
      to: null,
      start: null,
      end: null,
      endOpen: false,
      halfDate: [],
      datesHalf: null
    };

    this.saveEdit = this.saveEdit.bind(this);
    this.handleOnChange = this.handleOnChange.bind(this);
    this.handleChangeTypeOfLeave = this.handleChangeTypeOfLeave.bind(this);
    this.disabledDateBack = this.disabledDateBack.bind(this);
  }

  componentWillMount() {
    console.log(" ----------------- Request-Leave-Edit ----------------- ");
  }

  componentDidMount() {
    if (
      localStorage.getItem("role") !== "employee" &&
      localStorage.getItem("role") !== "supervisor"
    ) {
      this.props.history.push("/");
    }

    if (this.props.history.location.state === undefined) {
      this.props.history.push("/");
    } else {
      let id = Number(this.props.history.location.pathname.split("/").pop());
      let leave = this.props.history.location.state.leave.filter(
        el => el.id === id
      );
      this.props.fetchedEdit(leave[0]);
      this.props.typeLeaveFetchData();
    }
  }

  componentWillReceiveProps(nextProps) {
    if (nextProps.leave !== this.props.leave) {
      this.setState({
        start: this.formatDate(moment(nextProps.leave.date_from, "DD-MM-YYYY")),
        end: this.formatDate(moment(nextProps.leave.date_to, "DD-MM-YYYY")),
        datesHalf: nextProps.leave.half_dates.toString().split(",")
      });
    }
  }

  saveEdit = e => {
    e.preventDefault();
    this.props.saveEditLeave(this.props.leave, url => {
      this.props.history.push(url);
    });
  };

  handleOnChange = e => {
    let newLeave = {
      ...this.props.leave,
      [e.target.name]: e.target.value
    };
    this.props.handleEdit(newLeave);
  };

  handleChangeTypeOfLeave(value) {
    if (value === "11" || value === "44" || value === "55") {
      let typeLeave = {
        ...this.props.leave,
        type_leave_id: Number(value),
        reason: ""
      };
      this.props.handleEdit(typeLeave);
    } else {
      let typeLeave = {
        ...this.props.leave,
        type_leave_id: Number(value)
      };
      this.props.handleEdit(typeLeave);
    }
  }

  handleChangeSelect(value) {
    console.log("selected=======>", value);
  }

  disabledStartDate = startValue => {
    const endValue = this.state.to;
    if (!startValue || !endValue) {
      return false;
    }
    return startValue.valueOf() > endValue.valueOf();
  };

  disabledEndDate = endValue => {
    const startValue = this.state.from;
    if (!endValue || !startValue) {
      return false;
    }
    return endValue.valueOf() <= startValue.valueOf();
  };

  onChange = (field, value) => {
    this.setState({
      [field]: value
    });
  };

  onStartChange = value => {
    if (value !== null) {
      const date = new Date(value._d),
        mnth = ("0" + (date.getMonth() + 1)).slice(-2),
        day = ("0" + date.getDate()).slice(-2);
      let newDate = [day, mnth, date.getFullYear()].join("-");
      let newStart = [mnth, day, date.getFullYear()].join("-");

      let dateFrom = {
        ...this.props.leave,
        date_from: newDate
      };
      this.props.handleEdit(dateFrom);

      this.onChange("start", newStart);
    }
    this.onChange("from", value);
  };

  onEndChange = value => {
    if (value !== null) {
      const date = new Date(value._d),
        mnth = ("0" + (date.getMonth() + 1)).slice(-2),
        day = ("0" + date.getDate()).slice(-2);
      let newDate = [day, mnth, date.getFullYear()].join("-");
      let newEnd = [mnth, day, date.getFullYear()].join("-");

      let dateTo = {
        ...this.props.leave,
        date_to: newDate
      };
      this.props.handleEdit(dateTo);

      this.onChange("end", newEnd);
    }
    this.onChange("to", value);
  };

  disabledDate(current) {
    return current && current < moment().startOf("day");
  }

  disabledDateSick(current) {
    return (
      current &&
      current <
        moment()
          .subtract(7, "days")
          .startOf("day")
    );
  }

  disabledDateBack(current) {
    return this.state.to && this.state.to > current;
  }

  formatDate(value) {
    const date = new Date(value._d),
      mnth = ("0" + (date.getMonth() + 1)).slice(-2),
      day = ("0" + date.getDate()).slice(-2);
    return [mnth, day, date.getFullYear()].join("-");
  }

  handleStartOpenChange = open => {
    if (!open) {
      this.setState({ endOpen: true });
    }
  };

  handleEndOpenChange = open => {
    this.setState({ endOpen: open });
  };

  getDates(start, end) {
    var startDate = new Date(start);
    var endDate = new Date(end);
    let dates = [];
    while (startDate <= endDate) {
      var weekDay = startDate.getDay();
      if (weekDay < 6 && weekDay > 0) {
        var month = startDate.getMonth() + 1;
        if (month <= 9) {
          month = "0" + month;
        }
        var day = startDate.getDate();
        if (day <= 9) {
          day = "0" + day;
        }
        dates.push(day + "-" + month + "-" + startDate.getFullYear());
      }
      startDate.setDate(startDate.getDate() + 1);
    }
    return dates;
  }

  onChangeAddHalfDay(e) {
    let hiddenDiv = document.getElementById("halfDay");
    if (e.target.checked === true) {
      hiddenDiv.style.display = "block";
    } else {
      hiddenDiv.style.display = "none";
    }
    console.log(`checked add hald day = ${e.target.checked}`);
  }

  onChangeIsHalfDay(e, value) {
    console.log(`${e.target.value} checked is ${e.target.checked}`);
    if (e.target.checked) {
      this.setState(prevState => ({
        halfDate: update(prevState.halfDate, { $push: [e.target.value] })
      }));
    } else {
      let array = this.state.halfDate;
      let index = array.indexOf(e.target.value);
      this.setState(prevState => ({
        halfDate: update(prevState.halfDate, { $splice: [[index, 1]] })
      }));
    }

    // let halfDay = {
    //   ...this.props.leave,
    //   half_dates: this.state.halfDate
    // };
    // this.props.handleEdit(halfDay);
    // console.log("halfday==========>", halfDay);
  }

  onBackOn = value => {
    if (value !== null) {
      const date = new Date(value._d),
        mnth = ("0" + (date.getMonth() + 1)).slice(-2),
        day = ("0" + date.getDate()).slice(-2);
      let newDate = [day, mnth, date.getFullYear()].join("-");

      let backOn = {
        ...this.props.leave,
        back_on: newDate,
        half_dates: this.state.halfDate
      };
      this.props.handleEdit(backOn);
    }
  };

  handleBlur() {
    console.log("blur");
  }

  handleFocus() {
    console.log("focus");
  }

  render() {
    const { start, end, endOpen, datesHalf } = this.state;
    const formItemLayout = {
      labelCol: {
        xs: { span: 24 },
        sm: { span: 8 }
      },
      wrapperCol: {
        xs: { span: 24 },
        sm: { span: 16 }
      },
      style: {}
    };

    const formStyle = {
      width: "100%"
    };

    const elements = [];
    const dates = this.getDates(start, end);
    const dateFormat = "DD-MM-YYYY";
    let typeID;

    if (datesHalf !== null) {
      for (let i in dates) {
        for (let j in datesHalf) {
          if (datesHalf[j].trim() === dates[i]) {
            elements.push(
              <Checkbox
                key={j}
                id="is_half_day"
                name="is_half_day"
                value={datesHalf[j]}
                defaultChecked={true}
                onChange={e => this.onChangeIsHalfDay(e, datesHalf[j])}
              >
                {datesHalf[j]}
              </Checkbox>,
              <br />
            );
          } else {
            elements.push(
              <Checkbox
                key={i}
                id="is_half_day"
                name="is_half_day"
                onChange={e => this.onChangeIsHalfDay(e, dates[i])}
                value={dates[i]}
              >
                {dates[i]}
              </Checkbox>,
              <br />
            );
          }
        }
      }
    }

    this.props.typeLeave.map(d => {
      if (d.type_name === this.props.leave.type_name) {
        typeID = d.id;
      }
      return d;
    });

    if (this.props.loading) {
      return <Loading />;
    }
    return (
      <Layout>
        <HeaderNav />
        <Content
          className="container"
          style={{
            display: "flex",
            margin: "24px 16px 0",
            justifyContent: "center",
            paddingBottom: "73px"
          }}
        >
          <div
            style={{
              padding: 150,
              paddingBottom: 50,
              paddingTop: 50,
              background: "#fff",
              minHeight: 360
            }}
          >
            <h1> Form Edit Leave Request </h1>
            <Form onSubmit={this.handleSubmit} className="login-form">
              <FormItem {...formItemLayout} label="Type Of Leave">
                <Select
                  id="type_leave_id"
                  name="type_leave_id"
                  optionFilterProp="children"
                  onChange={this.handleChangeTypeOfLeave}
                  onSelect={(value, event) =>
                    this.handleChangeSelect(value, event)
                  }
                  defaultValue={this.props.leave.type_name}
                  // value={this.props.leave.type_id}
                  style={formStyle}
                  onFocus={this.handleFocus}
                  onBlur={this.handleBlur}
                >
                  {this.props.typeLeave.map(d => (
                    <Option key={d.id} value={d.id}>
                      {d.type_name}
                    </Option>
                  ))}
                </Select>
              </FormItem>

              {this.props.leave.type_leave_id === 22 ||
              this.props.leave.type_leave_id === 33 ||
              this.props.leave.type_leave_id === 66 ? (
                <FormItem {...formItemLayout} label="Reason">
                  <Input
                    type="text"
                    id="reason"
                    name="reason"
                    placeholder="reason"
                    onChange={this.handleOnChange}
                    style={formStyle}
                  />
                </FormItem>
              ) : (
                ""
              )}

              {this.props.leave.type_leave_id === 22 ||
              this.props.leave.type_leave_id === 33 ? (
                <FormItem {...formItemLayout} label="From">
                  <DatePicker
                    id="date_from"
                    name="date_from"
                    format={dateFormat}
                    defaultValue={moment(
                      this.props.leave.date_from,
                      "DD-MM-YYYY"
                    )}
                    onChange={this.onStartChange}
                    onOpenChange={this.handleStartOpenChange}
                    disabledDate={this.disabledDateSick}
                    style={formStyle}
                  />
                </FormItem>
              ) : (
                <FormItem {...formItemLayout} label="From">
                  <DatePicker
                    id="date_from"
                    name="date_from"
                    format={dateFormat}
                    defaultValue={moment(
                      this.props.leave.date_from,
                      "DD-MM-YYYY"
                    )}
                    onChange={this.onStartChange}
                    onOpenChange={this.handleStartOpenChange}
                    disabledDate={this.disabledDate}
                    style={formStyle}
                  />
                </FormItem>
              )}

              <FormItem {...formItemLayout} label="To">
                <DatePicker
                  id="date_to"
                  name="date_to"
                  format={dateFormat}
                  defaultValue={moment(this.props.leave.date_to, "DD-MM-YYYY")}
                  onChange={this.onEndChange}
                  open={endOpen}
                  onOpenChange={this.handleEndOpenChange}
                  disabledDate={this.disabledEndDate}
                  style={formStyle}
                />
              </FormItem>

              <FormItem>
                <Checkbox
                  id="add_half_day"
                  name="add_half_day"
                  onChange={this.onChangeAddHalfDay}
                  style={formStyle}
                >
                  Add Half Day
                </Checkbox>
              </FormItem>

              <div id="halfDay">
                <FormItem {...formItemLayout} label="Half Day">
                  {elements}
                </FormItem>
              </div>

              <FormItem {...formItemLayout} label="Back to work on">
                <DatePicker
                  id="back_on"
                  name="back_on"
                  format={dateFormat}
                  placeholder="Back to work"
                  defaultValue={moment(this.props.leave.back_on, "DD-MM-YYYY")}
                  disabledDate={this.disabledDateBack}
                  onChange={this.onBackOn}
                  style={formStyle}
                />
              </FormItem>

              <FormItem {...formItemLayout} label="Contact Address">
                <TextArea
                  type="text"
                  id="contact_address"
                  name="contact_address"
                  placeholder="contact_address, email, etc"
                  value={this.props.leave.contact_address}
                  onChange={this.handleOnChange}
                  autosize={{ minRows: 2, maxRows: 8 }}
                  style={formStyle}
                />
              </FormItem>

              <FormItem {...formItemLayout} label="Contact Number">
                <Input
                  type="text"
                  id="contact_number"
                  name="contact_number"
                  placeholder="Phone number"
                  value={this.props.leave.contact_number}
                  onChange={this.handleOnChange}
                  style={formStyle}
                />
              </FormItem>

              <FormItem>
                <Button
                  onClick={this.saveEdit}
                  htmlType="submit"
                  type="primary"
                  style={{
                    width: "35%"
                  }}
                >
                  EDIT
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
  loading: state.editRequestReducer.loading,
  leave: state.editRequestReducer.leave,
  typeLeave: state.fetchTypeLeaveReducer.typeLeave
});

const WrappedRegisterForm = Form.create()(LeaveEditPage);

const mapDispatchToProps = dispatch =>
  bindActionCreators(
    {
      fetchedEdit,
      handleEdit,
      saveEditLeave,
      typeLeaveFetchData
    },
    dispatch
  );
export default connect(
  mapStateToProps,
  mapDispatchToProps
)(WrappedRegisterForm);
