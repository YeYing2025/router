import React from 'react';
import { useParams } from 'react-router-dom';
import GroupsManager from '../../components/GroupsManager';
import { AppSection } from '../../router-ui';

const Group = () => {
  const { id: detailGroupId } = useParams();
  return (
    <div className='dashboard-container'>
      <AppSection>
        <GroupsManager detailGroupId={detailGroupId || ''} />
      </AppSection>
    </div>
  );
};

export default Group;
