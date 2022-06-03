

const Layout = ({children, props}) => (
    <div className="container" {...props} 
        style={{marginTop: '10rem', display: 'Flex', backgroundColor: '', padding: '10rem', justifyContent: 'center'}}
        >
        {children}
    </div>
)

export default Layout
