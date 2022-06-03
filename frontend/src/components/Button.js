
const AppButton = ({action, name}) => (
    <button className="btn waves-effect waves-light" type="submit" name="action" onClick={action}>
        {name}
        <i className="material-icons right">send</i>
    </button>
)

export default AppButton