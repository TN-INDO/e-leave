import React from "react";
import { Layout } from "antd";
const { Footer } = Layout;

export default class FooterLayout extends React.Component {
  render() {
    return (
      <Footer className="App-footer">
        <p>
          <a href="http://www.tnis.com">PT. TNIS Service Indonesia</a> &copy; 2018. All
          Right Reserved.
        </p>
      </Footer>
    );
  }
}
