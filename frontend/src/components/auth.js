import React from 'react';
import cookie from 'react-cookie';
import axios from 'axios';

export class Auth extends React.Component {

    constructor(props) {
        super(props);


        const uuidV4 = require('uuid/v4');

        this.state = {
            authToken: "",
            secureToken: uuidV4()
        };
    }
    
    componentDidMount = ()=>{
      
         axios.get("http://localhost/api/authtoken?token=" + this.state.secureToken)
             .then(res => {
                 console.log(res)
                 const auth = res.data;
                 this.setState({ authToken: auth });
               });
     }
    
    render() {
        console.log(this.state.authToken)
        
        console.log(cookie.load('postmurumsession'))
        
                if(this.state.authToken.trim() === "") {
                    return (
                        <div>
                        
                        <a href={'http://localhost/api/login?token=' + this.state.secureToken}>Please loginasdf</a>
                        </div>
                    )
                }
                else {
                    return (
                        <div>
                        
                        Welcome back {this.state.authToken}
                        </div>
                    )
                }            
            
 
        
    }
}
  
export default Auth