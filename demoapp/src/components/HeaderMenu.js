import React from 'react';
import { Menu } from 'semantic-ui-react';
import { Link } from 'react-router-dom';

class HeaderMenu extends React.Component{
    state={}
    handleItemClick = (e, { name }) => this.setState({ activeItem: name })
    render(){
        const { activeItem } = this.state
        return(
            <div>
                <Menu color='blue' stackable fixed='bottom' size='huge' borderless>
                    <Menu.Item header as={Link} to="/" active={activeItem === ''} 
                    onClick={this.handleItemClick}> DateChat 
                    </Menu.Item>
                    <Menu.Item position='right' size='large' as={Link} to="/signup"
                        name='Sign Up'
                        active={activeItem === 'Sign Up'}
                        onClick={this.handleItemClick}/>
                    <Menu.Item  as={Link} to="/login"
                        name='Login'
                        active={activeItem === 'Login'}
                        onClick={this.handleItemClick}/>
                    <Menu.Item as={Link} to="/about"
                        name='About us'
                        active={activeItem === 'aboutUs'}
                        onClick={this.handleItemClick}/>
                </Menu>
            </div>
        )
    }
}

export default HeaderMenu;