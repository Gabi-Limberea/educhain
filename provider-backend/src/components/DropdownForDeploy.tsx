import Dropdown from 'react-bootstrap/Dropdown';
import DropdownButton from 'react-bootstrap/DropdownButton';

function DropdownForDeploy() {
	return (
        <DropdownButton id="dropdown-basic-button" variant='secondary' title="Generate diplomas">
			<Dropdown.Item href="#/action-1">For all students</Dropdown.Item>
			<Dropdown.Item href="#/action-2">For students</Dropdown.Item>
			<Dropdown.Item href="#/action-3">Something else</Dropdown.Item>
        </DropdownButton>
    );
}

export default DropdownForDeploy;