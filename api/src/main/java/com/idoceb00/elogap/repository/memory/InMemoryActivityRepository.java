package com.idoceb00.elogap.repository.memory;

import com.idoceb00.elogap.domain.Activity;
import com.idoceb00.elogap.domain.MatchResult;
import com.idoceb00.elogap.repository.ActivityRepository;
import org.springframework.stereotype.Repository;

import java.time.Duration;
import java.time.Instant;
import java.util.List;
import java.util.Optional;

@Repository
public class InMemoryActivityRepository implements ActivityRepository {

    private final List<Activity> data;

    public InMemoryActivityRepository() {
        Instant now = Instant.now();

        this.data = List.of(
                new Activity(
                        "match_1",
                        "player_1",
                        MatchResult.win,
                        "RANKED_SOLO",
                        "Ahri",
                        "MID",
                        10, 2, 8,
                        210,
                        1714,
                        24500,
                        32,
                        now.minus(Duration.ofHours(2))
                ),
                new Activity(
                        "match_2",
                        "player_1",
                        MatchResult.loss,
                        "RANKED_SOLO",
                        "Jinx",
                        "ADC",
                        4, 7, 6,
                        260,
                        1980,
                        30200,
                        18,
                        now.minus(Duration.ofHours(6))
                )
        );
    }

    @Override
    public List<Activity> list() {
        return data;
    }

    @Override
    public Optional<Activity> findById(String id) {
        return data.stream().filter(a -> a.id().equals(id)).findFirst();
    }
}