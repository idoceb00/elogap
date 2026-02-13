package com.idoceb00.elogap.repository;

import com.idoceb00.elogap.domain.Activity;

import java.util.List;
import java.util.Optional;

public interface ActivityRepository {
    List<Activity> list();
    Optional<Activity> findById(String id);
}
