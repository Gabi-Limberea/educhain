import 'bootstrap-icons/font/bootstrap-icons.css';
import '../static/account.css'
import avatarPlaceholder from '../resources/avatar-placeholder.png';

function Account() {

    return (
        <div className="account-grid">
            <div className="top-container">
                <img className='profile-picture' src={avatarPlaceholder} alt="profile picture placeholder"></img>
                <div className='profile-header-text'> Provider's name </div>
            </div>

            <div className="bottom-row">
                <div className="left-container container-border">
                    <div className='user-info-top'> 
                        <h4 className='user-info-text'>User details</h4>
                        <button className='btn btn-secondary user-info-button'> Edit profile </button>
                    </div>
                    <ul className='user-info-list'>
                        <li> 
                            <dt>Email address</dt>
                            <dl>ceva@yahoo.com</dl> 
                        </li>
                        <li>  
                            <dt>Address</dt>
                            <dl>Iuliu Maniu nr 142</dl>
                        </li>
                        <li>
                            <dt>Contact phone </dt>
                            <dl>+40123456789</dl>
                        </li>
                    </ul>
                </div>

                <div className="right-container">
                    <div className="right-top container-border">
                        <div className='user-info-top'> 
                            <h4 className='user-info-text'>Certifications</h4>
                        </div>
                        <ul className='user-info-list' id='certification-list'>
                            
                        </ul>
                    </div>

                    <div className="right-bottom container-border">Other info</div>
                </div>
            </div>
        </div>
    )
} 

export default Account;