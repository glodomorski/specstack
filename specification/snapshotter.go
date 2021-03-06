package specification

type SnapshotterMatchFunc func(*Scenario) bool

type Snapshotter struct {
	ReadSourcer  ReadSourcer
	ObjectHasher ObjectHasher
	MatchFuncs   []SnapshotterMatchFunc
}

func NewSnapshotter(
	readSourcer ReadSourcer,
	objectHasher ObjectHasher,
	matchFuncs ...SnapshotterMatchFunc) *Snapshotter {
	return &Snapshotter{
		ReadSourcer:  readSourcer,
		ObjectHasher: objectHasher,
		MatchFuncs:   matchFuncs,
	}
}

func (s *Snapshotter) Snapshot(spec *Specification) (Snapshot, error) {
	q := NewQuery(spec).MapReduce(
		MapScenarios(),
		MapScenarioFileOrder(),
	)
	for _, fn := range s.MatchFuncs {
		q.MapReduce(MapScenarioMatchFunc(fn))
	}

	snapshot := Snapshot{}

	scenarios, err := s.snapshotScenarios(q.Scenarios())
	if err != nil {
		return snapshot, err
	}
	snapshot.Scenarios = scenarios

	return snapshot, nil
}

func (s *Snapshotter) snapshotScenarios(scenarios []*Scenario) ([]ScenarioSnapshot, error) {
	ss := make([]ScenarioSnapshot, len(scenarios))
	for i, scenario := range scenarios {
		storyID, err := s.DeterministicID(scenario.Story)
		if err != nil {
			return nil, err
		}
		scenarioID, err := s.DeterministicID(scenario)
		if err != nil {
			return nil, err
		}
		ss[i] = ScenarioSnapshot{
			StorySource: scenario.Story.Source(),
			StoryID:     storyID,
			ScenarioID:  scenarioID,
			LineNumber:  scenario.Location.Line,
		}
	}
	return ss, nil
}

func (s *Snapshotter) DeterministicID(object Sourcer) (string, error) {
	reader, err := s.ReadSourcer.ReadSource(object)
	if err != nil {
		return "", err
	}
	return s.ObjectHasher.ObjectHash(reader)
}
