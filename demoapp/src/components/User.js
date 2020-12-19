import React from 'react';
import {Component} from 'react';
import {Segment} from 'semantic-ui-react';


class User extends Component{
    render(){
        return(
            <div>
                <Segment>
        <p>Username : {this.props.email}</p>
                </Segment>
            </div>
        )
    }
}

export default User;