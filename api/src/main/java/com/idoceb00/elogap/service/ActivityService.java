package com.idoceb00.elogap.service;

import com.idoceb00.elogap.domain.Activity;
import com.idoceb00.elogap.domain.MatchResult;
import com.idoceb00.elogap.repository.ActivityRepository;
import org.springframework.stereotype.Service;

import java.util.List;
import java.util.Optional;
import java.util.stream.Stream;

@Service
public class ActivityService {
    private final ActivityRepository repository;

    public ActivityService(ActivityRepository repository) {
        this.repository = repository;
    }

    public List<Activity> list(Optional<MatchResult> result, Optional<String> champion) {
        Stream<Activity> stream = repository.list().stream();

        if (result.isPresent()) {
            stream = stream.filter(a -> a.result() == result.get());
        }

        if (champion.isPresent()) {
            stream = stream.filter(a -> a.champion().equalsIgnoreCase(champion.get()));
        }

        return stream.toList();
    }

    public Optional<Activity> findById(String id) {
        return repository.findById(id);
    }
}
