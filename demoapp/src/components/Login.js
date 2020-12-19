import React from 'react';
import { Form, Grid, Button,Segment } from 'semantic-ui-react';
import axios from 'axios';
//import { Menu } from 'semantic-ui-react';

class Login extends React.Component{
    state={
        'email' : '',
        'password' : '',
    }

    handleChange = (event) => {
        this.setState({[event.target.name]: event.target.value})
      }
    onSubmit = (e) => {
        e.preventDefault();
        console.log(this.state);
        axios.post('http://localhost:8081/login',this.state).then(function (response) {
            console.log(response);
          })
          .catch(function (error) {
            console.log(error);
          });
        //this.props.history.push('/user');

    }

    render(){
        return(
            <div >
                <p></p>
            <Grid centered width={12}>
                <Grid.Column style={{ maxWidth: 320 }}>
                <Segment textAlign='center' color='olive'>
                    <h4>Login</h4>
                </Segment>
                <Segment color='blue'>
                <Form size='tiny' id='login' name='login' color='blue' onSubmit={this.onSubmit}>
                        <Form.Input label='Email' placeholder='joe@schmoe.com' name="email" onChange={this.handleChange}/>
                        <Form.Input type='Password' label='Password' placeholder='Password' name="password" onChange={this.handleChange}/> 
                        <Button content='Login' primary size='tiny' />
                </Form>
                </Segment>
                </Grid.Column>
            </Grid>
            </div>
        )
    }
}

export default Login;