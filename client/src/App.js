import React, { Component } from "react";
import Modal from "./components/Modal";
import axios from 'axios';  

// Create a mock task data array
const taskData = [
  {
    id: 1,
    title: "Task 1",
    description: "Description for Task 1",
    completed: false
  },
  {
    id: 2,
    title: "Task 2",
    description: "Description for Task 2",
    completed: true
  },
  // Add more task objects as needed
];

class App extends Component {
  constructor(props) {
    super(props);
    this.state = {
      viewCompleted: false,
      activeItem: {
        title: "",
        description: "",
        completed: false
      },
      taskList: []
    };
  }

  // Add componentDidMount()
  componentDidMount() {
    this.refreshList();
  }

 
  refreshList = () => {
    axios   //Axios to send and receive HTTP requests
      .get("http://localhost:8081/tasks")
      .then(res => this.setState({ taskList: taskData }))
      // .then(res => console.log(res))
      .catch(err => console.log(err));
  };


  displayCompleted = status => {
    if (status) {
      return this.setState({ viewCompleted: true });
    }
    return this.setState({ viewCompleted: false });
  };


  renderTabList = () => {
    return (
      <div className="my-5 tab-list">
        <span
          onClick={() => this.displayCompleted(true)}
          className={this.state.viewCompleted ? "active" : ""}
        >
          completed
            </span>
        <span
          onClick={() => this.displayCompleted(false)}
          className={this.state.viewCompleted ? "" : "active"}
        >
          Incompleted
            </span>
      </div>
    );
  };

  // Main variable to render items on the screen
  renderItems = () => {
    const { viewCompleted } = this.state;
    const newItems = this.state.taskList.filter(
      item => item.completed === viewCompleted
    );
    return newItems.map(item => (
      <li
        key={item.id}
        className="list-group-item d-flex justify-content-between align-items-center"
      >
        <span
          className={`todo-title mr-2 ${this.state.viewCompleted ? "completed-todo" : ""
            }`}
          title={item.description}
        >
          {item.title}
        </span>
        <span>
          <button
            onClick={() => this.editItem(item)}
            className="btn btn-secondary mr-2"
          >
            Edit
          </button>
          <button
            onClick={() => this.handleDelete(item)}
            className="btn btn-danger"
          >
            Delete
          </button>
        </span>
      </li>
    ));
  };
  // ///////////////////////////////////////////////////////////

  ////add this after modal creation
  toggle = () => {//add this after modal creation
    this.setState({ modal: !this.state.modal });//add this after modal creation
  };
  // handleSubmit = item => {//add this after modal creation
  //   this.toggle();//add this after modal creation
  //   alert("save" + JSON.stringify(item));//add this after modal creation
  // };

  // Submit an item
  handleSubmit = item => {
    console.log("Priniting before ...");
    console.log(item);
    console.log("Priniting after ...");
    this.toggle();
    if (item.id) {
      // if old post to edit and submit
      axios
        .post(`http://localhost:8081/edit`, item)
        // .then(res => this.refreshList());
        .then(res => console.log(res));
      return;
    }
    // if new post to submit
    axios
      .post("http://localhost:8081/add", item)
      // .then(res => this.refreshList());
      .then(res => console.log(res));
  };

  // Delete item
  handleDelete = item => {
    axios
      .post(`http://localhost:8081/delete`)
      // .then(res => this.refreshList());
      .then(res => console.log(res));
  };
  // handleDelete = item => {//add this after modal creation
  //   alert("delete" + JSON.stringify(item));//add this after modal creation
  // };

  // Create item
  createItem = () => {
    const item = { title: "", description: "", completed: false };
    this.setState({ activeItem: item, modal: !this.state.modal });
  };

  //Edit item
  editItem = item => {
    this.setState({ activeItem: item, modal: !this.state.modal });
  };


  // -I- Start by visual effects to viewer
  render() {
    return (
      <main className="content">
        <h1 className="text-black text-uppercase text-center my-4">Task Manager</h1>
        <div className="row ">
          <div className="col-md-6 col-sm-10 mx-auto p-0">
            <div className="card p-3">
              <div className="">
                <button onClick={this.createItem} className="btn btn-primary">
                  Add task
                </button>
              </div>
              {this.renderTabList()}
              <ul className="list-group list-group-flush">
                {this.renderItems()}
              </ul>
            </div>
          </div>
        </div>
        {this.state.modal ? (
          <Modal
            activeItem={this.state.activeItem}
            toggle={this.toggle}
            onSave={this.handleSubmit}
          />
        ) : null}
      </main>
    );
  }
}
export default App;