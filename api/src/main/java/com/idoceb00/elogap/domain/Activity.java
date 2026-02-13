package com.idoceb00.elogap.domain;

import java.time.Instant;

public record Activity(
    String id,
    String playerId,
    MatchResult result,
    String queueType,
    String champion,
    String role,
    int kills,
    int deaths,
    int assists,
    int cs,
    int duration, // seconds
    int damage,
    int vision,
    Instant playedAt

) {}
