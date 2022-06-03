import AppButton from "./Button"

const SendButton = ({ sendMessage}) => (
    <div>
        <AppButton action={sendMessage} name={"Send Message"}/>
    </div>
)

export default SendButton