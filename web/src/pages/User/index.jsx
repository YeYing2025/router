import React from 'react';
import UsersTable from '../../components/UsersTable';
import { AppSection } from '../../router-ui';

const User = () => {
  return (
    <div className='dashboard-container'>
      <AppSection>
        <UsersTable />
      </AppSection>
    </div>
  );
};

export default User;
