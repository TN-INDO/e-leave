import React, { Component } from "react";
export default class NotFound extends Component {
  render() {
    return (
      <div className="message">
        <p className="rotate">ERROR</p>
        <h1 className="error">404</h1>
        <a className="btn" href="/">
          BACK TO HOME
        </a>
      </div>
    );
  }
}
