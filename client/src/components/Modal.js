import React, { Component } from "react";
// importing all of these classes from reactstrap module
import {
  Button,
  Modal,
  ModalHeader,
  ModalBody,
  ModalFooter,
  Form,
  FormGroup,
  Input,
  Label
} from "reactstrap";

class CustomModel extends Component {

    constructor(props){
        super(props);
        this.state = {
            activeItem: this.props.activeItem
        }
    }

    handleChange = e => {
        let { name, value} = e.target;
        if( e.target.type === "checkbox" ){
            value = e.target.checked;
        }
        const activeItem = { ...this.state.activeItem, [name]: value };
        this.setState({ activeItem })
    }

}