import AppButton from "./Button";


const LogIn = ({handleChange, login}) => (
    <div>
       <div className="center">
        <h4> Log in</h4>
       </div>
        <div>
            <div className="row" style={{backgroundColor: '', width: '30rem'}}>
                <form className="col s12" >
                    <div className="row">
                        <div className="input-field col s12">
                        <input id="username" type="text" className="validate" name="username" onChange={handleChange}/>
                        <label for="username">User Name</label>
                        </div>
                    </div>
                    <div className="row"> 
                        <div className="input-field col s12">
                        <input id="password" type="password" className="validate" name="password" onChange={handleChange}/>
                        <label for="password">Password</label>
                        </div>
                    </div>
                    <div className="row"> 
                        <div className="input-field col s12">
                            <AppButton name={"Login"} action={login}/>
                        </div>
                    </div>                    
                </form>
            </div>   
        </div>
    </div>
)

export default LogIn;