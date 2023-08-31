import logo from './logo.svg';
import React, {Component} from 'react';
import './App.css';

const tasks = [
  {
    id: 1,
    title: "Order Release",
    description: "Checkout customers and release",
    completed: true
  },
  {
    id: 2,
    title: "Restock Shelves",
    description: "Refill shelves with products",
    completed: false
  },
  {
    id: 3,
    title: "Customer Support",
    description: "Assist customers with inquiries",
    completed: false
  },
  {
    id: 4,
    title: "Inventory Check",
    description: "Count and verify inventory",
    completed: true
  }
];

// Define a class component named App
class App extends Component {

  // Constructor method - called when an instance of the component is created
  constructor(props) {
    // Call the constructor of the parent class (Component)
    super(props);

    // Define the initial state of the component
    this.state = {
      // Set the initial value of viewCompleted property in state to false
      viewCompleted: false,
      // Set the initial value of taskList property in state to an array of tasks (presumably defined elsewhere)
      taskList: tasks
    };
  }

  // A method to control the display of completed tasks
  displayCompleted = status => {
    // If the provided status is truthy
    if (status) {
      // Update the state using the setStatus method (should be corrected to setState)
      return this.setStatus({ viewCompleted: true });
    }
    // If the provided status is falsy
    return this.setStatus({ viewCompleted: false });
  }

  // This function returns a JSX element that represents a set of tabs for toggling between completed and incomplete tasks.
  renderTabList = () => {
    return (
      <div className='my-5 tab-list'>
        {/* Tab for showing completed tasks */}
        <span
          onClick={() => this.displayCompleted(true)}  
          className={this.state.viewCompleted ? "active" : ""}  
        >
          Completed
        </span>

        {/* Tab for showing incomplete tasks */}
        <span
          onClick={() => this.displayCompleted(false)}  
          className={this.state.viewCompleted ? "" : "active"}  
        >
          Incomplete
        </span>
      </div>
    );
  }

  // Filters tasks based on completed status and stores them in newItems array.
  renderItems = () => {
    const { viewCompleted } = this.state;  // Get whether to show completed or incomplete tasks
    const newItems = this.state.taskList.filter(
      item => item.completed == viewCompleted  // Filter tasks based on viewCompleted
    );
    return newItems.map(item => (
      <li key={item.id} className='list-group-item d-flex justify-content-between align-items-center'>
        <span className={`todo-title mr-2 ${this.state.viewCompleted ? "completed-todo" : ""}`} title={item.title} >
          {item.title}
        </span>
        <span>
          <button className='btn btn-info mr-2'>Edit</button>
          <button className='btn btn-info mr-2'>Delete</button>
        </span>
      </li>
    ))
  };

  render() {
    return (
      <main className='context'>
        <h1 className='text-black text-uppercase text-center my-4'>ConcurTask</h1>
        <h1 className='row'>
          <div className='col-md-6 col-sma-10 mx-auto p-0'>
            <div className='card p-3'>
              <div>
                <button className='btn btn-warning'>Add Task</button>
              </div>
              {this.renderTabList()}
              <ul className='list-group list-group-flush'>
                {this.renderItems()}
              </ul>
            </div>
          </div>
        </h1>
      </main>
    )
  } 
}

export default App;