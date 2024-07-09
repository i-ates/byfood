import React, { useState, useEffect } from "react";
import axios from "axios";
import Button from "react-bootstrap/Button";
import Table from "react-bootstrap/Table";
import "./UserListView.css";
import { EditUserScreen } from "./EditUserScreen";

export function UserListView() {
    const [users, setUsers] = useState([]); // State to store user data
    const [selectedUser, setSelectedUser] = useState(null); // State to store the selected user for editing
    const [showModal, setShowModal] = useState(false); // State to control the display of the edit user modal

    // Function to toggle the edit user modal and set the selected user
    const toggleModal = (user) => {
        setSelectedUser(user);
        setShowModal((prevShowModal) => !prevShowModal);
    };

    // Function to reload the main view
    const reloadMainView = () => {
        window.location.reload();
    };

    // Function to delete a user
    const deleteUser = async (user) => {
        try {
            const response = await axios.post(
                `${process.env.REACT_APP_API_URL}/api/delete-user`,
                user
            );
            console.log("User deleted:", response.data);
        } catch (error) {
            console.error("Error while deleting user:", error);
        }
        setTimeout(() => {
            reloadMainView();
        }, 500);
    };

    // Fetch users from the API on component mount
    useEffect(() => {
        axios
            .get(`${process.env.REACT_APP_API_URL}/api/get-users`)
            .then((resp) => setUsers(resp.data))
            .catch((error) => console.error("Error fetching users:", error));
    }, []);

    return (
        <div className="user-list-container">
            <Table striped bordered hover className="custom-table">
                <thead>
                <tr>
                    <th>ID</th>
                    <th>Name</th>
                    <th>Email</th>
                    <th>Age</th>
                </tr>
                </thead>
                <tbody>
                {users !== null ? (
                    users.map((user) => (
                        <tr key={user.id}>
                            <td>{user.id}</td>
                            <td>{user.name}</td>
                            <td>{user.mail}</td>
                            <td>{user.age}</td>
                            <td>
                                <Button
                                    variant="primary"
                                    onClick={() => toggleModal(user)}
                                    className="custom-edit-button"
                                >
                                    Edit
                                </Button>
                            </td>
                            <td>
                                <Button
                                    variant="primary"
                                    onClick={() => deleteUser(user)}
                                    className="custom-edit-button"
                                >
                                    Delete
                                </Button>
                            </td>
                        </tr>
                    ))
                ) : (
                    <tr>
                        <td colSpan="6">No user exists...</td>
                    </tr>
                )}
                </tbody>
            </Table>
            {/* Render the edit user modal if showModal is true and a user is selected */}
            {showModal && selectedUser !== null && (
                <EditUserScreen
                    show={showModal}
                    onHide={() => toggleModal(null)}
                    user={selectedUser}
                />
            )}
        </div>
    );
}
