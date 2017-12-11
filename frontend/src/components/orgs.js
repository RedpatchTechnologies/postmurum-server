import React from 'react';
import axios from 'axios';

class Organization extends React.Component {

    constructor(props) {
        super(props);

        this.state = {
          posts: []
        };
    }


   componentDidMount = ()=>{
      
        axios.get("http://localhost/api/org")
            .then(res => {
                console.log(res)
                const posts = res.data;
                this.setState({ posts });
              });
    }


    render() {
        return (
          <div>
            
            <ul>
              {this.state.posts.map(post =>
                <li key={post.id}>{post.name}</li>
              )}
            </ul>
          </div>
        );
      }
}


export default Organization