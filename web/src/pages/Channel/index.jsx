import React from 'react';
import ChannelsTable from '../../components/ChannelsTable';
import { AppSection } from '../../router-ui';

const Channel = () => {
  return (
    <div className='dashboard-container'>
      <AppSection>
        <ChannelsTable />
      </AppSection>
    </div>
  );
};

export default Channel;
