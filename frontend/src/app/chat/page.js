"use client";
import React, { Component } from "react";
// import "./App.css";
import { connect, sendMsg } from "../api/page";

class Chat extends Component {
  constructor(props) {
    super(props);
    connect();
  }

  send() {
    console.log("hello");
    sendMsg("hello");
  }

  render() {
    return (
      <div className="Chat">
        <button onClick={this.send}>Hit</button>
      </div>
    );
  }
}

export default Chat;
