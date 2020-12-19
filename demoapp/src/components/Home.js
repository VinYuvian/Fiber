import React from 'react';
import { Grid, Segment,Header,Icon } from 'semantic-ui-react';

class Home extends React.Component{
    render(){
        return(
            <div>
                <p></p>
                <Grid centered verticalAlign='bottom'>
                    <Grid.Row>
                    <Grid.Column style={{ maxWidth: 640 }} textAlign='center'>  
                    <p></p>                     
                    <Segment>
                        <h2>Welcome to DateChat</h2>
                    </Segment>
                    <p>Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, 
                        when an unknown printer took a galley of type and scrambled it to make a type 
                        specimen book. It has survived not only five centuries, but also the leap into
                        electronic typesetting, remaining essentially unchanged. It was popularised in 
                        the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, 
                        and more recently with desktop publishing software like Aldus PageMaker 
                        including versions of Lorem Ipsum.Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, 
                        when an unknown printer took a galley of type and scrambled it to make a type 
                        specimen book.</p>
                    </Grid.Column>
                    </Grid.Row>
                </Grid>
                <Grid columns={3} padded='horizontally'>
                <Grid.Row>
                    <Grid.Column style={{ maxWidth: 640 }} textAlign='center'>
                        <Segment padded placeholder>
                        <Header icon>
                            <Icon name='search' />
                             Find Country
                        </Header>
                        </Segment>
                    </Grid.Column>
                    <Grid.Column style={{ maxWidth: 640 }} textAlign='center'>
                        <Segment placeholder>
                        <Header icon>
                            <Icon name='users' />
                             Make a Friiend
                        </Header>
                        </Segment>
                    </Grid.Column>
                    <Grid.Column style={{ maxWidth: 640 }} textAlign='center'>
                        <Segment placeholder>
                        <Header icon>
                            <Icon name='search' />
                             Find Country
                        </Header>
                        </Segment>
                    </Grid.Column>
                </Grid.Row>
                </Grid>
            </div>
        )
    }
}

export default Home;