import Home from './components/home';
import Time from './components/time';
import Counter from './components/counter';
import './app.css'

const { useState, useEffect } = React;

export default function App() {
    const [route, setRoute] = useState(window.location.pathname);

    useEffect(() => {
        const handlePopState = () => setRoute(window.location.pathname);

        window.addEventListener('popstate', handlePopState);

        return () => window.removeEventListener('popstate', handlePopState);
    }, []);

    const navigate = (path) => {
        window.history.pushState({}, '', path);
        setRoute(path);
    };

    let Component;
    switch (route) {
        case '/time':
            Component = Time;
            break;
        case '/counter':
            Component = Counter;
            break;
        default:
            Component = Home;
    }

    return (
        <>
            <nav style={{ backgroundColor: '#f0f0f0', padding: '10px', textAlign: 'center' }}>
                <button onClick={() => navigate('/')} style={{ margin: '0 10px', padding: '5px 10px' }}>Home</button>
                <button onClick={() => navigate('/time')} style={{ margin: '0 10px', padding: '5px 10px' }}>Time</button>
                <button onClick={() => navigate('/counter')} style={{ margin: '0 10px', padding: '5px 10px' }}>Counter</button>
            </nav>
            <Component />
        </>)
}

ReactDOM.render(<App />, document.getElementById('root'));
