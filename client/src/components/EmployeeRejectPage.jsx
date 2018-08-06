import React, { Component } from "react";
import { bindActionCreators } from "redux";
import { connect } from "react-redux";
import {
  employeeGetRequestReject,
  employeeDeleteRequest
} from "../store/Actions/employeeAction";
import {
  Layout,
  Table,
  Modal,
  Button,
  Input,
  Icon,
  Divider,
  Popconfirm,
  message
} from "antd";
import HeaderNav from "./menu/HeaderNav";
import Loading from "./menu/Loading";
import Footer from "./menu/Footer";
const { Content } = Layout;
let data;

class EmployeeRejectPage extends Component {
  constructor(props) {
    super(props);
    this.state = {
      loading: false,
      visible: false,
      user: null,
      data: this.props.leaves,
      filterDropdownVisible: false,
      filtered: false,
      searchID: ""
    };
  }

  componentWillMount() {
    console.log(
      "------------ Employee-List-Reject-Request -------------------"
    );
  }

  componentWillReceiveProps(nextProps) {
    if (nextProps.leaves !== this.props.leaves) {
      this.setState({ data: nextProps.leaves });
    }
    data = nextProps.leaves;
  }

  componentDidMount() {
    if (!localStorage.getItem("token")) {
      this.props.history.push("/");
    } else if (
      localStorage.getItem("role") !== "employee" &&
      localStorage.getItem("role") !== "supervisor"
    ) {
      this.props.history.push("/");
    }
    this.props.employeeGetRequestReject();
  }

  onSearchID = () => {
    const { searchID } = this.state;
    const reg = new RegExp(searchID, "gi");
    this.setState({
      filterDropdownIDVisible: false,
      filtered: !!searchID,
      data: data
        .map(record => {
          const match = String(record.id).match(reg);
          if (!match) {
            return null;
          }
          return {
            ...record,
            ID: (
              <span>
                {this.state.data.map(
                  (text, i) =>
                    String(text.id) === searchID ? (
                      <span key={i} className="highlight">
                        {text}
                      </span>
                    ) : (
                      text
                    ) // eslint-disable-line
                )}
                }
              </span>
            )
          };
        })
        .filter(record => !!record)
    });
  };

  onInputChangeID = e => {
    this.setState({
      searchID: e.target.value
    });
  };

  showDetail = record => {
    this.setState({
      visible: true,
      user: record
    });
  };

  employeeDeleteRequest = (leaves, id) => {
    this.props.employeeDeleteRequest(leaves, id);
  };

  onSelectChange = selectedRowKeys => {
    console.log("selected row: ", selectedRowKeys);
  };

  handleCancel = () => {
    this.setState({ visible: false });
  };

  onShowSizeChange(current, pageSize) {
    console.log(current, pageSize);
  }

  render() {
    const { visible, loading } = this.state;
    const columns = [
      {
        title: "ID",
        dataIndex: "id",
        key: "id",
        width: 90,
        filterDropdown: (
          <div className="custom-filter-dropdown-id">
            <Input
              type="number"
              ref={ele => (this.searchInput = ele)}
              placeholder="Search request id"
              value={this.state.searchID}
              onChange={this.onInputChangeID}
              onPressEnter={this.onSearchID}
            />
            <Button type="primary" onClick={this.onSearchID}>
              Search
            </Button>
          </div>
        ),
        filterIcon: (
          <Icon
            type="search"
            style={{ color: this.state.filtered ? "#108ee9" : "#aaa" }}
          />
        ),
        filterDropdownIDVisible: this.state.filterDropdownIDVisible,
        onFilterDropdownVisibleChange: visible => {
          this.setState(
            {
              filterDropdownIDVisible: visible
            },
            () => this.searchInput && this.searchInput.focus()
          );
        }
      },
      {
        title: "Employee Number",
        dataIndex: "employee_number",
        key: "employee_number",
        width: 100
      },
      {
        title: "Name",
        dataIndex: "name",
        key: "name",
        width: 150
      },
      {
        title: "Position",
        dataIndex: "position",
        key: "position",
        width: 150
      },
      {
        title: "Email",
        dataIndex: "email",
        key: "email",
        width: 150
      },
      {
        title: "Type Of Leave",
        dataIndex: "type_name",
        key: "type_name",
        width: 150
      },
      {
        title: "From",
        dataIndex: "date_from",
        key: "date_from",
        width: 120
      },
      {
        title: "To",
        dataIndex: "date_to",
        key: "date_to",
        width: 120
      },
      {
        title: "Action",
        key: "action",
        width: 200,
        render: (value, record) => (
          <span>
            <Button type="primary" onClick={() => this.showDetail(record)}>
              Detail
            </Button>
            <Divider type="vertical" />
            <Popconfirm
              placement="top"
              title={"Are you sure delete this leave request?"}
              onConfirm={() => {
                this.employeeDeleteRequest(this.props.leaves, record.id);
                message.success("Leave request has been delete!");
              }}
              okText="Yes"
              cancelText="No"
            >
              <Button type="danger">Delete</Button>
            </Popconfirm>
          </span>
        )
      }
    ];

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
              paddingBottom: "606px"
            }}
          >
            <div style={{ padding: 20, background: "#fff" }}>
              <Table
                columns={columns}
                dataSource={this.state.data}
                rowKey={record => record.id}
                onRowClick={this.onSelectChange}
                pagination={{
                  className: "my-pagination",
                  defaultCurrent: 1,
                  defaultPageSize: 5,
                  total: 50,
                  showSizeChanger: this.onShowSizeChange
                }}
              />
            </div>

            <Modal
              visible={visible}
              title="Detail Leave Request Rejected"
              onCancel={this.handleCancel}
              style={{ top: "20" }}
              bodyStyle={{ padding: "0" }}
              footer={[
                <Button
                  key="cancel"
                  loading={loading}
                  onClick={this.handleCancel}
                >
                  Return
                </Button>
              ]}
            >
              <div style={{ padding: 10, background: "#fff" }}>
                ID : {this.state.user && this.state.user.id} <br />
                Name : {this.state.user && this.state.user.name} <br />
                Gender : {this.state.user && this.state.user.gender} <br />
                Email : {this.state.user && this.state.user.email} <br />
                Type Of Leave : {this.state.user &&
                  this.state.user.type_name}{" "}
                <br />
                Reason : {this.state.user && this.state.user.reason} <br />
                From : {this.state.user && this.state.user.date_from} <br />
                To : {this.state.user && this.state.user.date_to} <br />
                Half Day : {this.state.user && this.state.user.half_dates !== "" ? ( this.state.user.half_dates ):("none")} <br />
                Back On : {this.state.user && this.state.user.back_on} <br />
                Total Leave : {this.state.user &&
                  this.state.user.total} day <br />
                Leave Balance :{" "}
                {this.state.user && this.state.user.leave_remaining} day <br />
                Contact Address :{" "}
                {this.state.user && this.state.user.contact_address} <br />
                Contact Number :{" "}
                {this.state.user && this.state.user.contact_number} <br />
                Status : {this.state.user && this.state.user.status} <br />
                Reject Reason :{" "}
                {this.state.user && this.state.user.reject_reason}
              </div>
            </Modal>
          </Content>
          <Footer />
        </Layout>
      );
    }
  }
}

const mapStateToProps = state => ({
  loading: state.fetchEmployeeReducer.loading,
  leaves: state.fetchEmployeeReducer.leaves
});

const mapDispatchToProps = dispatch =>
  bindActionCreators(
    {
      employeeGetRequestReject,
      employeeDeleteRequest
    },
    dispatch
  );

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(EmployeeRejectPage);
