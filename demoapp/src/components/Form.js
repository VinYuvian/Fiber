import React from 'react';
import { Form, Grid, Button,Segment } from 'semantic-ui-react';
import axios from 'axios';
//import { route } from 'react-router-dom';

const options = [
    { key: 'm', text: 'Male', value: 'male' },
    { key: 'f', text: 'Female', value: 'female' },
    { key: 'o', text: 'Other', value: 'other' },
  ]

class LoginForm extends React.Component{
    state={
        firstname : '',
        lastname : '',
        email : '',
        password : '',
        gender : ''
    }

    handleChange = (event) => {
        this.setState({[event.target.name]: event.target.value})
      }
    
    handleDropdownChange = (event, data) => {
        this.setState({[data.name]:data.value})
    }
    onSubmit = (e) => {
        e.preventDefault();
        console.log(this.state);
        axios.post('http://localhost:8081/signup',this.state).then(function (response) {
            console.log(response.data);
          })
          .catch(function (error) {
            console.log(error);
          });
        //this.props.history.push('/login');
    }

    render(){
        return(
            <div >
            
                <p></p>
            <Grid centered width={12}>
                <Grid.Column style={{ maxWidth: 320 }}>
                <Segment textAlign='center' color='olive'>
                <h4>Sign Up</h4>
                </Segment>
                <Segment color='blue'>
                <Form size='tiny' color='blue' onSubmit={this.onSubmit}>
                        <Form.Input label='First Name' placeholder='First Name' name="firstname" onChange={this.handleChange}/>           
                        <Form.Input label='Last Name' placeholder='Last Name' name="lastname" onChange={this.handleChange}/>
                        <Form.Input label='Email' placeholder='joe@schmoe.com' name="email" onChange={this.handleChange}/>
                        <Form.Input type='Password' label='Password' placeholder='Password' name="password" onChange={this.handleChange}/>
                        <Form.Select options={ options } placeholder='Gender' name="gender" onChange={this.handleDropdownChange} />
                        <Form.Checkbox label='I agree to the Terms and Conditions' name="termState"/> 
                        <Button content='Sign Up' primary size='tiny' />
                </Form>
                </Segment>
                </Grid.Column>
            </Grid>
            </div>
        )
    }
}

export default LoginForm;